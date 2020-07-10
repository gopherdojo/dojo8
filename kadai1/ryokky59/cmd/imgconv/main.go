package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gopherdojo/dojo8/kadai1/ryokky59/convert"
)

var converter convert.ImgConverter

func init() {
	flag.StringVar(&converter.Dir, "dir", "./", "対象のディレクトリ")
	flag.StringVar(&converter.From, "from", "jpg", "変換する対象の画像の拡張子")
	flag.StringVar(&converter.To, "to", "png", "変換した後の画像の拡張子")
	flag.Parse()
}

func main() {
	if converter.From != "png" &&
		converter.From != "jpeg" &&
		converter.From != "jpg" &&
		converter.From != "gif" &&
		converter.From != "bmp" &&
		converter.From != "tiff" {
		fmt.Errorf("Unsupported extension in from")
		os.Exit(1)
	}

	if converter.To != "png" &&
		converter.To != "jpeg" &&
		converter.To != "jpg" &&
		converter.To != "gif" &&
		converter.To != "bmp" &&
		converter.To != "tiff" {
		fmt.Errorf("Unsupported extension in to")
		os.Exit(1)
	}

	if converter.From == converter.To {
		fmt.Errorf("from and to are conflict")
		os.Exit(1)
	}

	if f, err := os.Stat(converter.Dir); os.IsNotExist(err) || !f.IsDir() {
		fmt.Errorf("Directory does not exist")
		os.Exit(1)
	}

	if err := convert.Exec(&converter); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
