/*
main は画像変換のmainパッケージです。

実行時のコマンドライン引数は下記になります。

  -ft:
     変換する画像の形式（gif, jpeg, jpg, png, bmp, tiff）
  -fdir:
     変換する画像の保存しているディレクトリ
  -tt:
     変換後の画像の形式（gif, jpeg, jpg, png, bmp, tiff）
  -tdir:
     変換後の画像を保存するディレクトリ
*/
package main

import (
	"bubu/cmd/converter"
	"flag"
	"fmt"
	"os"
)

// main は画像変換のメイン処理となり、コマンドライン引数の値をConverterに渡して、画像変換を実行します。
func main() {
	var (
		fromType = flag.String("ft", "jpeg", "変換する画像の形式 (gif, jpeg, jpg, png, bmp, tiff)")
		fromDir  = flag.String("fdir", "", "変換する画像の保存しているディレクトリ")
		toType   = flag.String("tt", "png", "変換後の画像の形式 (gif, jpeg, jpg, png, bmp, tiff)")
		toDir    = flag.String("tdir", "", "変換後の画像を保存するディレクトリ")
	)
	flag.Parse()

	cnv := new(converter.Converter)
	cnv.SetDir(*fromDir, *toDir)
	err := cnv.Convert(*fromType, *toType)

	if err != nil {
		str := fmt.Sprintf("err: %s", err)
		fmt.Println(str)
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
