package main

import (
	"flag"
	"github.com/gopherdojo/dojo8/kadai1/tsuchinaga/conv"
	"log"
)

func main() {
	var dir, src, dest string
	flag.StringVar(&dir, "dir", "", "変換する画像のあるディレクトリ")
	flag.StringVar(&src, "src", "jpeg", "optional 変換元の画像形式 jpeg|png")
	flag.StringVar(&dest, "dest", "png", "optional 変換後の画像形式 jpeg|png")
	flag.Parse()

	// validation
	if dir == "" {
		log.Fatalln("dirの指定は必須です")
	}
	if !conv.IsDir(dir) {
		log.Fatalf("%sは存在しないかディレクトリではありません\n", dir)
	}
	if !conv.IsValidFileType(src) {
		log.Fatalf("%sは許可されていない画像形式です\n", src)
	}
	if !conv.IsValidFileType(dest) {
		log.Fatalf("%sは許可されていない画像形式です", dest)
	}
	if src == dest {
		log.Fatalln("srcとdestで違う画像形式を選択してください")
	}

	// 変換実行
	conv.ExecConvert(dir, src, dest)
}
