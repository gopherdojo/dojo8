//画像変換を行います。
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gopherdojo/dojo8/kadai1/InAyk/converter"
)

//オプション指定
var bf = flag.String("bf", "jpg", "Image format before conversion")
var af = flag.String("af", "png", "Image format after conversion")
var dir = flag.String("dir", "", "Directory containing images to be converted")

func main() {
	//オプションの解析
	flag.Parse()

	//入力されたディレクトリのチェックをかける
	if err := checkDir(); err != nil {
		//入力チェックで返ってきたらエラーを出力し、プログラムを終了させる
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	convert()

}

//checkDir 指定されたディレクトリのチェックを行います。
//チェック内容：引数の値、ディレクトリの存在、ディレクトリを表しているか
func checkDir() error {
	//引数の数チェック
	if *dir == "" {
		return fmt.Errorf("No directory specified")
	}

	//ディレクトリの存在チェック ＆　ディレクトリを表しているかどうか
	if m, err := os.Stat(*dir); os.IsNotExist(err) {
		return err
	} else if !m.IsDir() {
		return fmt.Errorf("%s : not a directory", *dir)
	}

	return nil

}

//convert 指定ディレクトリに対し、画像変換を行います。
func convert() {

	ca := converter.Args{
		Bf: strings.ToLower(*bf),
		Af: strings.ToLower(*af),
	}

	//ディレクトリを再帰的にチェック
	filepath.Walk(*dir, func(path string, info os.FileInfo, err error) error {

		//対象の画像形式のみ画像変換ImageConverter()の実行
		if strings.ToLower(filepath.Ext(path)) == "."+*bf {
			ca.FilePath = path
			if err := converter.ImageConverter(ca); err != nil {
				//エラーが返ってきても、エラーを表示しプログラムは続行
				fmt.Fprintln(os.Stderr, err)
			}
		}
		return nil
	})

}
