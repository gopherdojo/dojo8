package main

import (
	"flag"
	"github.com/gopherdojo/dojo8/kadai1/tsuchinaga/imgconv"
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
	if isDir, err := imgconv.IsDir(dir); err != nil || !isDir {
		log.Fatalf("%sは存在しないかディレクトリではありません\n", dir)
	}
	if !imgconv.IsValidFileType(src) {
		log.Fatalf("%sは許可されていない画像形式です\n", src)
	}
	if !imgconv.IsValidFileType(dest) {
		log.Fatalf("%sは許可されていない画像形式です", dest)
	}
	if src == dest {
		log.Fatalln("srcとdestで違う画像形式を選択してください")
	}

	// 変換実行
	imgconv.Do(dir, src, dest)
}
