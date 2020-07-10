package imageconverter

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

type Args struct {
	Directory       string
	BeforeExtension string
	AfterExtension  string
}

func exec(source, dest string) (err error) {
	sourceFile, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("file could not be opened. target: %s", source)
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("image file could not be created. target: %s", dest)
	}

	defer func() {
		if err == nil {
			err = destFile.Close()
		}
	}()

	img, _, err := image.Decode(sourceFile)
	if err != nil {
		return err
	}

	switch strings.ToLower(filepath.Ext(dest)) {
	case ".png":
		err = png.Encode(destFile, img)
	case ".jpg", ".jpeg":
		err = jpeg.Encode(destFile, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	case ".gif":
		err = gif.Encode(destFile, img, &gif.Options{NumColors: 256})
	case ".bmp":
		err = bmp.Encode(destFile, img)
	case ".tiff", ".tif":
		err = tiff.Encode(destFile, img, nil)
	default:
		err = fmt.Errorf("image file could not be created due to an unknown extension. target: %s", dest)
	}

	return err
}

type FilepathWalk struct {
	Walk func(root string, walkFn filepath.WalkFunc) error
}

func NewFilepathWalk() FilepathWalk {
	return FilepathWalk{Walk: filepath.Walk}
}

// 指定したディレクトリ以下のJPGファイルをPNGに変換（デフォルト）
// ディレクトリ以下は再帰的に処理する
// 変換前と変換後の画像形式を指定できる（オプション）
func Convert(args Args, filepathWalk FilepathWalk) error {
	return filepathWalk.Walk(args.Directory, func(path string, info os.FileInfo, err error) error {
		return convertImage(path, info, err, args)
	})
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
	return exec(path, replaceExt(info.Name(), "."+args.AfterExtension))
}

func replaceExt(filePath, toExt string) string {
	return filePath[:len(filePath)-len(filepath.Ext(filePath))] + toExt
}
