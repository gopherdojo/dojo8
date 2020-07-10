package convert

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

// Exec 画像の変換を実行する
func Exec(converter *ImgConverter) error {
	err := filepath.Walk(converter.Dir, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(info.Name(), converter.From) {
			converter.Name = strings.TrimRight(info.Name(), converter.From)
			err := converter.convert()
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// ImgConverter 画像の拡張子を変換する
type ImgConverter struct {
	Dir  string
	Name string
	From string
	To   string
}

func (c *ImgConverter) convert() error {
	file, err := os.Open(c.path())
	if err != nil {
		return err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	dst, err := os.Create(c.dstPath())
	defer dst.Close()
	if err != nil {
		return err
	}

	switch strings.ToLower(c.To) {
	case "png":
		err = png.Encode(dst, img)
		if err != nil {
			return err
		}
	case "jpeg", "jpg":
		err = jpeg.Encode(dst, img, nil)
		if err != nil {
			return err
		}
	case "gif":
		err = gif.Encode(dst, img, nil)
		if err != nil {
			return err
		}
	case "bmp":
		err = bmp.Encode(dst, img)
		if err != nil {
			return err
		}
	case "tiff":
		err = tiff.Encode(dst, img, nil)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("Unsupported extension")
	}

	return nil
}

func (c *ImgConverter) path() string {
	return c.Dir + c.Name + c.From
}

func (c *ImgConverter) dstPath() string {
	return c.Dir + c.Name + c.To
}
