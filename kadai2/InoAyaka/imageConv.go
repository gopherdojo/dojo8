//画像変換を行います。
package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/gopherdojo/dojo8/kadai2/InoAyaka/converter"
)

//オプション指定
var bf = flag.String("bf", "jpg", "Image format before conversion")
var af = flag.String("af", "png", "Image format after conversion")
var dir = flag.String("dir", "", "Directory containing images to be converted")

//テスト中は出力先を変更できるよう、パッケージレベルとする
var outErr io.Writer = os.Stderr

const (
	exitCodeSuccess = 0
	exitCodeErr     = 1
)

func main() {
	if errCode := run(); errCode == exitCodeErr {
		os.Exit(exitCodeErr)
	}
}

//
func run() int {
	//オプションの解析
	flag.Parse()

	//入力されたディレクトリのチェックをかける
	if err := checkDir(*dir); err != nil {
		fmt.Fprintln(outErr, err)
		return exitCodeErr
	}

	if err := convert(*bf, *af, *dir); err != nil {
		fmt.Fprintln(outErr, err)
		return exitCodeErr
	}

	return exitCodeSuccess
}

//checkDir 指定されたディレクトリのチェックを行います。
//チェック内容：引数の値、ディレクトリの存在、ディレクトリを表しているか
func checkDir(dir string) error {
	//引数の数チェック
	if dir == "" {
		return fmt.Errorf("No directory specified")
	}

	//ディレクトリの存在チェック ＆　ディレクトリを表しているかどうか
	if m, err := os.Stat(dir); os.IsNotExist(err) {
		return err
	} else if !m.IsDir() {
		return fmt.Errorf("%s : not a directory", dir)
	}

	return nil

}

//convert 指定ディレクトリに対し、画像変換を行います。
func convert(bf string, af string, dir string) error {

	ca := converter.Args{
		BeforeFormat: strings.ToLower(bf),
		AfterFormat:  strings.ToLower(af),
	}

	//ディレクトリを再帰的にチェック
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		//対象の画像形式のみ画像変換ImageConverter()の実行
		if strings.ToLower(filepath.Ext(path)) == "."+bf {
			ca.FilePath = path
			if err := converter.ImageConverter(ca); err != nil {
				//エラーが返ってきても、エラーを表示しプログラムは続行
				fmt.Fprintln(outErr, err)
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
