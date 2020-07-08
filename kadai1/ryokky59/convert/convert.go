
package convert

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"image"
	"image/png"
	"image/jpeg"
)

var (
	dir  string
	from string
	to   string
)

func init() {
	flag.StringVar(&dir, "dir", "./", "対象のディレクトリ")
	flag.StringVar(&from, "from", "jpg", "変換する対象の画像の拡張子")
	flag.StringVar(&to, "to", "png", "変換した後の画像の拡張子")
	flag.Parse()
}

// Exec 画像の変換を実行する
func Exec() error {
	converter, err := newImgConverter(dir, from, to)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = filepath.Walk(converter.dir, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(info.Name(), from) {
			converter.name = strings.TrimRight(info.Name(), from)
			err := converter.convert()
			if err != nil {
				fmt.Println(err)
				return nil
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return nil
}

type imgConverter struct {
	dir  string
	name string
	from string
	to   string
}

func newImgConverter(dir, from, to string) (*imgConverter, error) {
	if from != "png" && from != "jpeg" && from != "jpg" {
		return nil, fmt.Errorf("Unsupported extension in from")
	}

	if to != "png" && to != "jpeg" && to != "jpg" {
		return nil, fmt.Errorf("Unsupported extension in to")
	}

	if from == to {
		return nil, fmt.Errorf("from and to are conflict")
	}

	if f, err := os.Stat(dir); os.IsNotExist(err) || !f.IsDir() {
		return nil, fmt.Errorf("Directory does not exist")
	}

	return &imgConverter{
		dir: dir,
		from: from,
		to: to,
	}, nil
}

func (c *imgConverter) convert() error {
	file, err := os.Open(c.path())
	defer file.Close()
	if err != nil {
		return err
	}

	img, _, err :=  image.Decode(file)
	if err != nil {
		return err
	}

	dst, err := os.Create(c.dstPath())
	defer dst.Close()
	if err != nil {
		return err
	}

	switch c.to {
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
	}

	return nil
}

func (c *imgConverter) path() string {
	return c.dir + c.name + c.from
}

func (c *imgConverter) dstPath() string {
	return c.dir + c.name + c.to
}
