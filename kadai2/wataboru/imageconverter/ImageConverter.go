package imageconverter

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

var (
	convImage       func(path string, info os.FileInfo, err error, args Args) error
	execute         func(source, dest string) (err error)
	osOpen          func(name string) (*os.File, error)
	osCreate        func(name string) (*os.File, error)
	sourceFileClose func(file *os.File) error
	destFileClose   func(file *os.File) error
	imageDecode     func(r io.Reader) (image.Image, string, error)
	newImgEncoder   func(ext string) (Encoder, error)
)

func init() {
	convImage = convertImage
	execute = exec
	osOpen = os.Open
	sourceFileClose = func(file *os.File) error {
		return file.Close()
	}
	osCreate = os.Create
	destFileClose = func(file *os.File) error {
		return file.Close()
	}
	imageDecode = image.Decode
	newImgEncoder = newImageEncoder

}

type Args struct {
	Directory       string
	BeforeExtension string
	AfterExtension  string
}

type Encoder interface {
	Encode(w io.Writer, m image.Image) error
}

type JpegEncoder struct{}

func (enc JpegEncoder) Encode(w io.Writer, m image.Image) error {
	return jpeg.Encode(w, m, &jpeg.Options{Quality: jpeg.DefaultQuality})
}

type PngEncoder struct{}

func (enc PngEncoder) Encode(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}

type GifEncoder struct{}

func (enc GifEncoder) Encode(w io.Writer, m image.Image) error {
	return gif.Encode(w, m, &gif.Options{NumColors: 256})
}

type BmpEncoder struct{}

func (enc BmpEncoder) Encode(w io.Writer, m image.Image) error {
	return bmp.Encode(w, m)
}

type TiffEncoder struct{}

func (enc TiffEncoder) Encode(w io.Writer, m image.Image) error {
	return tiff.Encode(w, m, nil)
}

func newImageEncoder(dest string) (Encoder, error) {
	switch strings.ToLower(filepath.Ext(dest)) {
	case ".jpg", ".jpeg":
		return JpegEncoder{}, nil
	case ".png":
		return PngEncoder{}, nil
	case ".gif":
		return GifEncoder{}, nil
	case ".bmp":
		return BmpEncoder{}, nil
	case ".tiff", ".tif":
		return TiffEncoder{}, nil
	default:
		return nil, fmt.Errorf("image file could not be created due to an unknown extension. target: %s", dest)
	}
}

func exec(source, dest string) (err error) {
	sourceFile, err := osOpen(source)
	if err != nil {
		return fmt.Errorf("file could not be opened. target: %s", source)
	}
	defer sourceFileClose(sourceFile)

	destFile, err := osCreate(dest)
	if err != nil {
		return fmt.Errorf("image file could not be created. target: %s", dest)
	}

	defer func() {
		if err == nil {
			err = destFileClose(destFile)
		}
	}()

	img, _, err := imageDecode(sourceFile)
	if err != nil {
		return err
	}

	e, err := newImgEncoder(dest)
	if err != nil {
		return err
	}
	return e.Encode(destFile, img)
}

func convertImage(path string, info os.FileInfo, err error, args Args) error {
	if err != nil {
		return fmt.Errorf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
	}

	if info.IsDir() {
		return nil
	}

	ext := strings.ToLower(filepath.Ext(path))
	if "."+args.BeforeExtension != ext {
		return nil
	}
	return execute(path, replaceExt(info.Name(), "."+args.AfterExtension))
}

// 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
// ディレクトリ以下は再帰的に処理する
// 変換前と変換後の画像形式を指定できる（オプション）
func Convert(args Args) error {
	return filepath.Walk(args.Directory, func(path string, info os.FileInfo, err error) error {
		return convImage(path, info, err, args)
	})
}

func replaceExt(filePath, toExt string) string {
	return filePath[:len(filePath)-len(filepath.Ext(filePath))] + toExt
}
