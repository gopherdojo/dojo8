package imgconv

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path"

	"golang.org/x/xerrors"
)

var (
	ReadDirError               = xerrors.New("read dir error")
	FileStatError              = xerrors.New("file stat error")
	OpenFileError              = xerrors.New("open file error")
	CloseFileError             = xerrors.New("close file error")
	NotImageError              = xerrors.New("not image error")
	ReadImageError             = xerrors.New("read image error")
	CreateDestinationFileError = xerrors.New("create destination file error")
	EncodeImageError           = xerrors.New("encode image error")
)

// NewIMGConverter - 新しい画像変換の生成
func NewIMGConverter() IMGConverter {
	return &imgConverter{converter: NewConverter()} // converterを外からもらってもいいけど、使い方は決まってるので決め打ちで
}

// IMGConverter - 画像変換のユースケースのインターフェース
type IMGConverter interface {
	Do(path, src, dest string) chan error
}

// imgConverter - 画像変換のユースケース
type imgConverter struct {
	converter Converter
}

// Do - 指定されたディレクトリから再帰的にたどりながらファイルを変換する
func (c *imgConverter) Do(path, src, dest string) chan error {
	ch := make(chan error)

	go func() {
		defer close(ch)

		dirs, files, err := c.converter.DirFileList(path)
		if err != nil {
			ch <- err
			return
		}

		// 非同期で再帰的にディレクトリをたどる
		for _, dir := range dirs {
			dch := c.Do(dir, src, dest)
			for err := range dch {
				ch <- err
			}
		}

		// ファイルを変換する
		for _, file := range files {
			if err := c.converter.Convert(file, src, dest); err != nil {
				ch <- err
			}
		}
	}()

	return ch
}

// NewConverter - 新しいConverterを生成する
func NewConverter() Converter {
	return &converter{}
}

// converter - 変換処理とそれに付随する処理をもつサービスのインターフェース
type Converter interface {
	IsDir(path string) (bool, error)
	GetIMGType(path string) (_ string, err error)
	DirFileList(filePath string) ([]string, []string, error)
	Convert(filePath string, src string, dest string) (err error)
}

// converter - 変換処理とそれに付随する処理をもつサービス
type converter struct{}

// IsDir - pathがディレクトリかどうか
func (c converter) IsDir(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, xerrors.Errorf("%+v: %w", err, FileStatError)
	}
	return fi.IsDir(), nil
}

// GetIMGType - 引数のファイルの画像フォーマットを取得する
func (c converter) GetIMGType(path string) (_ string, err error) {
	f, err := os.Open(path)
	if err != nil { // 開けない
		return "", OpenFileError
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			err = xerrors.Errorf("%+v: %w", closeErr, CloseFileError)
		}
	}()

	_, format, err := image.DecodeConfig(f)
	if err != nil { // 画像じゃない
		return "", NotImageError
	}
	return format, nil
}

// DirFileList - 引数直下のディレクトリとファイルの配列を返す
func (c converter) DirFileList(filePath string) ([]string, []string, error) {
	infos, err := ioutil.ReadDir(filePath)
	if err != nil {
		return nil, nil, xerrors.Errorf("%+v(filePath: %s): %w", err, filePath, ReadDirError)
	}

	dirs, files := make([]string, 0), make([]string, 0)
	for _, info := range infos {
		jp := path.Join(filePath, info.Name())
		if info.IsDir() {
			dirs = append(dirs, jp)
		} else {
			files = append(files, jp)
		}
	}
	return dirs, files, nil
}

// Convert - 画像変換
func (c converter) Convert(filePath string, src string, dest string) (err error) {
	if ft, err := c.GetIMGType(filePath); err != nil {
		return err
	} else if ft != src {
		return nil
	}

	f, err := os.Open(filePath)
	if err != nil { // 開けない
		return xerrors.Errorf("%+v: %w", err, OpenFileError)
	}
	defer func() {
		if closeErr := f.Close(); closeErr != nil {
			err = xerrors.Errorf("%+v: %w", closeErr, CloseFileError)
		}
	}()

	img, _, err := image.Decode(f)
	if err != nil {
		return xerrors.Errorf("%+v: %w", err, ReadImageError)
	}

	newFilePath := fmt.Sprintf("%s.%s.%s", filePath, "converted", dest)
	o, err := os.Create(newFilePath)
	if err != nil {
		return xerrors.Errorf("%+v: %w", err, CreateDestinationFileError)
	}
	defer func() {
		if closeErr := o.Close(); closeErr != nil {
			err = xerrors.Errorf("%+v: %w", closeErr, CloseFileError)
		}
	}()

	switch dest {
	case "jpeg":
		if err = jpeg.Encode(o, img, nil); err != nil {
			return xerrors.Errorf("%+v: %w", err, EncodeImageError)
		}
	case "png":
		if err = png.Encode(o, img); err != nil {
			return xerrors.Errorf("%+v: %w", err, EncodeImageError)
		}
	}
	return nil
}
