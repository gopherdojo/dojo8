package main

import (
	"flag"
	"log"
	"os"

	"github.com/gopherdojo/dojo8/kadai2/tsuchinaga/imgconv"

	"github.com/gopherdojo/dojo8/kadai2/tsuchinaga/validation"
)

var (
	stderr       Logger = log.New(os.Stderr, "", log.LstdFlags|log.Lmicroseconds)
	validator           = validation.NewValidator()
	imgConverter        = imgconv.NewIMGConverter()
)

func main() {
	// 入力値受け取り
	var dir, src, dest string // ベースディレクトリ、変換元画像形式、変換後画像形式
	flag.StringVar(&dir, "dir", "", "変換する画像のあるディレクトリ")
	flag.StringVar(&src, "src", "jpeg", "optional 変換元の画像形式 jpeg|png")
	flag.StringVar(&dest, "dest", "png", "optional 変換後の画像形式 jpeg|png")
	flag.Parse()

	// validatorで入力値チェック
	switch {
	case !validator.IsValidDir(dir):
		stderr.Println("dirの指定は必須です")
		os.Exit(2)
	case !validator.IsValidSrc(src):
		stderr.Println("srcに有効な画像形式を指定してください")
		os.Exit(2)
	case !validator.IsValidDest(dest, src):
		stderr.Println("destに有効な画像形式を指定してください。また、srcとdestは一致しないよう指定してください")
		os.Exit(2)
	}

	// converterで変換の実行
	for err := range imgConverter.Do(dir, src, dest) {
		switch err {
		case imgconv.NotImageError: // 画像じゃないファイルも存在しうるので無視する
		default:
			stderr.Println(err)
		}

	}
}

// Logger - ログ出力のインターフェース
type Logger interface {
	Println(v ...interface{})
}
