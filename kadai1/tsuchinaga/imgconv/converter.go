package imgconv

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/xerrors"
)

var (
	ReadDirError               = xerrors.New("read dir error")
	FileStatError              = xerrors.New("file stat error")
	OpenFileError              = xerrors.New("open source file error")
	CreateDestinationFileError = xerrors.New("create destination file error")
	ReadImageError             = xerrors.New("read image error")
	EncodeImageError           = xerrors.New("encode image error")
	NoImageError               = xerrors.New("no image error")
)

var validFileTypes = map[string]bool{"jpeg": true, "png": true}

// IsValidFileType - 指定されたファイルタイプが利用可能かを返す
func IsValidFileType(fileType string) bool {
	return validFileTypes[fileType]
}

// IsDir - pathがディレクトリかどうか
func IsDir(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, xerrors.Errorf("%+v: %w", err, FileStatError)
	}
	return fi.IsDir(), nil
}

// Do - 変換の実行
func Do(dir, src, dest string) chan error {
	c := converter{
		dirList:      []string{dir},
		srcFileType:  src,
		destFileType: dest,
		ch:           make(chan error),
	}
	go func() {
		defer close(c.ch)
		err := c.exec()
		if err != nil {
			c.ch <- err
		}
	}()
	return c.ch
}

// converter - 変換機能の実装
type converter struct {
	dirList      []string
	srcFileType  string
	destFileType string
	ch           chan error
}

// exec - ディレクトリをたどりながら変換を実行
func (c *converter) exec() error {
	for len(c.dirList) > 0 {
		dirPath := c.dirList[0]
		c.dirList = c.dirList[1:]

		files, err := ioutil.ReadDir(dirPath)
		if err != nil {
			return xerrors.Errorf("%+v: %w", err, ReadDirError)
		}

		for _, file := range files {
			path := filepath.Join(dirPath, file.Name())
			if file.IsDir() {
				c.dirList = append(c.dirList, path)
			} else {
				if err := c.convert(path); err != nil {
					c.ch <- err // エラーをチャネルに渡して続きを進める
				}
			}
		}
	}
	return nil
}

// convert - 変換処理
func (c converter) convert(path string) (err error) {
	f, err := os.Open(path)
	if err != nil { // 開けない
		return xerrors.Errorf("%+v: %w", err, OpenFileError)
	}
	defer f.Close()

	ft, err := getFileType(path)
	if err != nil {
		return err
	}
	if ft == c.srcFileType {
		img, _, err := image.Decode(f)
		if err != nil {
			return xerrors.Errorf("%+v: %w", err, ReadImageError)
		}

		newFilePath := fmt.Sprintf("%s.%s", path, c.destFileType)
		o, err := os.Create(newFilePath)
		if err != nil {
			return xerrors.Errorf("%+v: %w", err, CreateDestinationFileError)
		}
		defer func() {
			err = o.Close()
		}()

		switch c.destFileType {
		case "jpeg":
			if err = jpeg.Encode(o, img, nil); err != nil {
				return xerrors.Errorf("%+v: %w", err, EncodeImageError)
			}
		case "png":
			if err = png.Encode(o, img); err != nil {
				return xerrors.Errorf("%+v: %w", err, EncodeImageError)
			}
		}
	}
	return nil
}

// getFileType - 画像ファイルの型を得る
func getFileType(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil { // 開けない
		return "", OpenFileError
	}
	defer f.Close()

	_, format, err := image.DecodeConfig(f)
	if err != nil { // 画像じゃない
		return "", NoImageError
	}
	return format, nil
}
