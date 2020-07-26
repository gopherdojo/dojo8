package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sync/errgroup"
)

type downloader struct {
	targetURL string
	dir       string
	div       uint
	fileName  string
	fileSize  uint
	split     uint
}

type Range struct {
	start  uint
	end    uint
	worker uint
}

var div = flag.Uint("div", 5, "Specifying the number of divisions")
var dir = flag.String("dir", "./", "Download directory")
var targetURL = flag.String("URL", "", "URL to download")

func main() {

	flag.Parse()

	//URLのチェック
	d := &downloader{
		div:       *div,
		dir:       *dir,
		targetURL: *targetURL,
	}

	if ok := d.ready(); !ok {
		os.Exit(1)
	}

	fmt.Println(strings.Repeat("-", 50))

	if err := d.download(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := d.unionFiles(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := d.deleteTmpFiles(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(strings.Repeat("-", 50))
	fmt.Println("download complete")

}

//download 分割ダウンロードを実行します
func (d *downloader) download() error {

	_, err := url.Parse(d.targetURL)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	grp, gctx := errgroup.WithContext(ctx)

	for i := 0; i < int(d.div); i++ {
		part := i + 1
		grp.Go(func() error {
			return d.partDownload(gctx, part)
		})
	}

	if err := grp.Wait(); err != nil {
		return err
	}

	return nil
}

func (d *downloader) partDownload(gctx context.Context, part int) error {

	r := d.makeRange(part)

	req, err := http.NewRequest("GET", d.targetURL, nil)
	if err != nil {
		return err
	}

	//ヘッダーのセット
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", r.start, r.end))

	//HTTP要求を送信
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to split get requests: %d", part)
	}
	defer resp.Body.Close()

	//分割ファイル名の設定
	tmpFilePath := fmt.Sprintf("%s/%s.%d.%d", d.dir, d.fileName, d.div, part)

	f, err := os.OpenFile(tmpFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("failed to create %s ", tmpFilePath)
	}
	defer f.Close()

	if _, err := io.Copy(f, resp.Body); err != nil {
		return err
	}

	//他のタスクからのキャンセルを受信した場合を考慮
	select {
	case <-gctx.Done():
		fmt.Printf("cancelled [%02d]\n", part)
		return gctx.Err()
	default:
		fmt.Printf("finished download [%02d]　: bytes  %5d - %5d\n", part, r.start, r.end)
		return nil
	}
}

//makeRange 範囲の指定を行うRangeを作成します
func (d *downloader) makeRange(part int) Range {
	start := d.split * uint(part-1)
	end := start + d.split - 1

	//最後のパートのみendの指定をファイルサイズに設定する
	if uint(part) == d.div {
		end = d.fileSize
	}

	return Range{
		start:  start,
		end:    end,
		worker: uint(part),
	}
}

//unionFiles ファイルの結合を行います
func (d *downloader) unionFiles() error {

	outFilePath := filepath.Join(d.dir, d.fileName)

	outFile, err := os.Create(outFilePath)
	if err != nil {
		return err
	}

	for i := 0; i < int(d.div); i++ {
		part := i + 1
		tmpFilePath := fmt.Sprintf("%s/%s.%d.%d", d.dir, d.fileName, d.div, part)

		tmpFile, err := os.Open(tmpFilePath)
		if err != nil {
			return fmt.Errorf("failed to open %s", tmpFilePath)
		}

		if _, err := io.Copy(outFile, tmpFile); err != nil {
			return fmt.Errorf("failed to copy %s", tmpFilePath)
		}

		if err := tmpFile.Close(); err != nil {
			return fmt.Errorf("failed to close %s", tmpFilePath)
		}

	}

	if err := outFile.Close(); err != nil {
		return err
	}

	return nil
}

//deleteTmpFiles 不要になった一時ファイルを削除します
func (d *downloader) deleteTmpFiles() error {
	for i := 0; i < int(d.div); i++ {
		part := i + 1
		tmpFilePath := fmt.Sprintf("%s/%s.%d.%d", d.dir, d.fileName, d.div, part)

		if err := os.Remove(tmpFilePath); err != nil {
			return fmt.Errorf("failed to remove %s", tmpFilePath)
		}
	}
	return nil
}
