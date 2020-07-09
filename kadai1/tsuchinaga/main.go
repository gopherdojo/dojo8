package main

import (
	"flag"
	"log"
	"os"

	"github.com/gopherdojo/dojo8/kadai1/tsuchinaga/imgconv"
)

func main() {
	var dir, src, dest string
	flag.StringVar(&dir, "dir", "", "変換する画像のあるディレクトリ")
	flag.StringVar(&src, "src", "jpeg", "optional 変換元の画像形式 jpeg|png")
	flag.StringVar(&dest, "dest", "png", "optional 変換後の画像形式 jpeg|png")
	flag.Parse()

	// validation
	if dir == "" {
		log.Println("dirの指定は必須です")
		os.Exit(2)
	}
	if isDir, err := imgconv.IsDir(dir); err != nil || !isDir {
		log.Printf("%sは存在しないかディレクトリではありません\n", dir)
		os.Exit(2)
	}
	if !imgconv.IsValidFileType(src) {
		log.Printf("%sは許可されていない画像形式です\n", src)
		os.Exit(2)
	}
	if !imgconv.IsValidFileType(dest) {
		log.Printf("%sは許可されていない画像形式です", dest)
		os.Exit(2)
	}
	if src == dest {
		log.Println("srcとdestで違う画像形式を選択してください")
		os.Exit(2)
	}

	// 変換実行
	if err := imgconv.Do(dir, src, dest); err != nil {
		log.Println(err)
		os.Exit(2)
	}
}
