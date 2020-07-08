package imgconv

import (
	"fmt"
	"golang.org/x/xerrors"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var (
	FileStatError = xerrors.New("file stat error")
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
func Do(dir, src, dest string) {
	c := converter{
		dirList:      []string{dir},
		srcFileType:  src,
		destFileType: dest,
	}
	c.exec()
}

// converter - 変換機能の実装
type converter struct {
	dirList      []string
	srcFileType  string
	destFileType string
}

// exec - ディレクトリをたどりながら変換を実行
func (c *converter) exec() {
	for len(c.dirList) > 0 {
		dirPath := c.dirList[0]
		c.dirList = c.dirList[1:]

		files, err := ioutil.ReadDir(dirPath)
		if err != nil {
			log.Printf("ディレクトリ: %sが読み込めなかったためスキップします\n", dirPath)
			continue
		}

		for _, file := range files {
			path := filepath.Join(dirPath, file.Name())
			if file.IsDir() {
				c.dirList = append(c.dirList, path)
			} else {
				c.convert(path)
			}
		}
	}
}

// convert - 変換処理
func (c converter) convert(path string) {
	f, err := os.Open(path)
	if err != nil { // 開けない
		log.Printf("ファイル: %sが開けなかったためスキップします\n", path)
		return
	}
	defer f.Close()

	if getFileType(path) == c.srcFileType {
		img, _, err := image.Decode(f)
		if err != nil {
			log.Printf("ファイル: %sが読み込めなかったためスキップします\n", path)
			return
		}

		newFilePath := fmt.Sprintf("%s.%s", path, c.destFileType)
		o, err := os.Create(newFilePath)
		if err != nil {
			log.Printf("変換後ファイル: %sが作成できなかったためスキップします\n", newFilePath)
			return
		}
		defer o.Close()

		err = nil
		switch c.destFileType {
		case "jpeg":
			err = jpeg.Encode(o, img, nil)

		case "png":
			err = png.Encode(o, img)
		}
		if err != nil {
			log.Printf("ファイル: %sの変換に失敗しました\n", path)
		} else {
			log.Printf("%s => %s\n", path, newFilePath)
		}
	}
}

// getFileType - 画像ファイルの型を得る
func getFileType(path string) string {
	f, err := os.Open(path)
	if err != nil { // 開けない
		return ""
	}
	defer f.Close()

	_, format, err := image.DecodeConfig(f)
	if err != nil { // 画像じゃない
		return ""
	}
	return format
}
