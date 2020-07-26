package downloader

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

type ProgressType string

const (
	ProgressTypeDone   = "done"
	ProgressTypeFailed = "failed"
	ProgressTypeNotice = "notice"
)

type Progress struct {
	Type    ProgressType
	Message string
}

// New - 新しいダウンローダーを生成する
func New(url string) Downloader {
	fileName := filepath.Base(url)
	return &downloader{url: url, fileName: fileName}
}

// Downloader - ダウンローダーのインターフェース
type Downloader interface {
	Do(context.Context) chan Progress
}

// downloader - ダウンローダー
type downloader struct {
	ch       chan Progress
	url      string
	fileName string
}

// Do - 分割しながらデータを取得する
func (d *downloader) Do(ctx context.Context) chan Progress {
	// すでに結果を返すchanが作られていたらそれを返す
	if d.ch != nil {
		return d.ch
	}

	d.ch = make(chan Progress)
	go func() {
		defer func() { d.ch <- Progress{Type: ProgressTypeDone} }() // 最後に必ずDoneが返される
		// 分割に対応しているかを確認する
		ranges, err := d.AcceptRanges()
		if err != nil {
			d.ch <- Progress{Type: ProgressTypeFailed, Message: err.Error()}
			return
		}

		totalSize, err := d.ContentLength()
		if err != nil {
			d.ch <- Progress{Type: ProgressTypeFailed, Message: err.Error()}
			return
		}
		if totalSize <= 0 {
			d.ch <- Progress{Type: ProgressTypeNotice, Message: "Content-Lengthが不明なため一括取得します"}
		}

		var bytesList [][]byte
		switch ranges {
		case "none": // 分割に対応していない場合はgoroutine1つで取得する
			if body, err := d.GetAll(ctx); err != nil {
				d.ch <- Progress{Type: ProgressTypeFailed, Message: err.Error()}
				return
			} else {
				bytesList = [][]byte{body}
				d.ch <- Progress{Type: ProgressTypeNotice, Message: "ダウンロード完了"}
			}
		case "bytes": // 分割に対応している場合、4分割にして4つのgoroutineで取得する
			var wg sync.WaitGroup
			bytesList = make([][]byte, 4)
			for i := 0; i < 4; i++ {
				n := i
				part := totalSize / 4
				wg.Add(1)
				go func() {
					defer wg.Done()
					start := part * int64(n)
					end := (part * int64(n+1)) - 1
					if n == 3 {
						end = totalSize
					}

					d.ch <- Progress{Type: ProgressTypeNotice, Message: fmt.Sprintf("分割ダウンロード(%d)開始", n+1)}
					if body, err := d.GetPart(ctx, start, end); err != nil {
						d.ch <- Progress{Type: ProgressTypeFailed, Message: err.Error()}
						return
					} else {
						bytesList[n] = body
						d.ch <- Progress{Type: ProgressTypeNotice, Message: fmt.Sprintf("分割ダウンロード(%d)完了", n+1)}
					}
				}()
			}
			wg.Wait()
		}

		out, err := os.Create(d.fileName)
		if err != nil {
			d.ch <- Progress{Type: ProgressTypeFailed, Message: err.Error()}
			return
		}
		for _, b := range bytesList {
			if _, err := out.Write(b); err != nil {
				d.ch <- Progress{Type: ProgressTypeFailed, Message: err.Error()}
				return
			}
		}
		if err := out.Close(); err != nil {
			d.ch <- Progress{Type: ProgressTypeFailed, Message: err.Error()}
			return
		}
	}()
	return d.ch
}

func (d *downloader) GetAll(ctx context.Context) (bytes []byte, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", d.url, nil)
	if err != nil {
		return nil, err
	}
	cli := new(http.Client)
	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (d *downloader) GetPart(ctx context.Context, start, end int64) (bytes []byte, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", d.url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Range", fmt.Sprintf("bytes=%d-%d", start, end))
	cli := new(http.Client)
	res, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// ContentLength - url先のコンテンツの長さを取得
func (d *downloader) ContentLength() (int64, error) {
	res, err := http.Head(d.url)
	if err != nil {
		return 0, err
	}
	return res.ContentLength, nil
}

// AcceptRanges - 範囲リクエストの設定を確認する
func (d *downloader) AcceptRanges() (string, error) {
	res, err := http.Head(d.url)
	if err != nil {
		return "", err
	}
	ars := res.Header["Accept-Ranges"]
	if len(ars) == 0 {
		return "none", nil
	} else {
		return ars[0], nil
	}
}
