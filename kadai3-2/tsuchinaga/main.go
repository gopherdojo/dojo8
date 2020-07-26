package main

import (
	"context"
	"flag"
	"github.com/gopherdojo/dojo8/kadai3-2/tsuchinaga/downloader"
	"log"
	"os"
	"os/signal"
)

func main() {
	flag.Parse()
	url := flag.Arg(0)

	ctx, cFunc := context.WithCancel(context.Background())

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	go func() {
		for range sigCh {
			println("すべての処理を中断して終了します")
			cFunc()
			close(sigCh)
			os.Exit(2)
		}
	}()

	dCh := downloader.New(url).Do(ctx)
	for progress := range dCh {
		switch progress.Type {
		case downloader.ProgressTypeFailed, downloader.ProgressTypeNotice:
			log.Println(progress.Message)
		case downloader.ProgressTypeDone:
			log.Println("完了")
			return
		default:
			log.Println("想定外のtype")
		}
	}
}
