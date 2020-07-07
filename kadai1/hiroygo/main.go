package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gopherdojo/dojo8/kadai1/hiroygo/imgconv"
)

func getSaveFilePath(srcFilePath string, t imgconv.ImageType) string {
	saveFileName := filepath.Base(srcFilePath)
	saveFileName = strings.TrimSuffix(saveFileName, filepath.Ext(saveFileName))
	saveFileName += "." + t.String()
	return filepath.Join(filepath.Dir(srcFilePath), saveFileName)
}

func parseArgs() (dir string, in, out imgconv.ImageType, err error) {
	d := flag.String("d", "./", "変換の対象となる画像が格納されたディレクトリパス")
	i := flag.String("i", "jpg", "変換の対象となる画像種別")
	o := flag.String("o", "png", "変換後の画像種別")
	flag.Parse()

	in = imgconv.ToImageType(*i)
	out = imgconv.ToImageType(*o)
	if in == imgconv.Unknown || out == imgconv.Unknown {
		in = imgconv.Unknown
		out = imgconv.Unknown
		err = fmt.Errorf("設定された画像種別が不正です")
		return
	}

	dir = *d
	return
}

func main() {
	dir, inType, outType, err := parseArgs()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	pathes, err := imgconv.GetImageFilePathesRecursive(dir, inType)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	for _, path := range pathes {
		m, err := imgconv.LoadImage(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s の読み込みに失敗しました。%v\n", path, err)
			continue
		}

		savePath := getSaveFilePath(path, outType)
		if err := imgconv.SaveImage(m, outType, savePath); err != nil {
			fmt.Fprintf(os.Stderr, "%s の保存に失敗しました。%v\n", savePath, err)
			continue
		}
	}
}
