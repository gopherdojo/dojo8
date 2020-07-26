package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func (d *downloader) ready() bool {

	errFlg := true

	if err := d.checkDir(); err != nil {
		errFlg = false
		fmt.Fprintf(os.Stderr, "checkDir(%s): NG \n", d.dir)
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Fprintf(os.Stdout, "checkDir(%s): OK\n", d.dir)
	}

	if err := d.checkURL(); err != nil {
		errFlg = false
		fmt.Fprintf(os.Stderr, "checkURL(%s): NG \n", d.targetURL)
		fmt.Fprintln(os.Stderr, err)
	} else {
		fmt.Fprintf(os.Stdout, "checkURL(%s): OK\n", d.targetURL)
	}

	d.setFileName()

	if err := d.setFileSize(); err != nil {
		errFlg = false
		fmt.Fprintf(os.Stderr, "getFileSize(%s): NG \n", d.targetURL)
		fmt.Fprintln(os.Stderr, err)
	}

	return errFlg
}

//checkURL 指定されたダウンロードURLのチェックを行います
func (d *downloader) checkURL() error {
	_, err := url.ParseRequestURI(d.targetURL)
	if err != nil {
		return fmt.Errorf("URL error , %w", err)
	}

	u, err := url.Parse(d.targetURL)
	if err != nil || u.Scheme == "" || u.Hostname() == "" {
		return fmt.Errorf("URL error , %w", err)
	}

	return nil
}

//checkDir 指定されたディレクトリのチェックを行います
//チェック内容：ディレクトリの存在、ディレクトリを表しているか
func (d *downloader) checkDir() error {

	//ディレクトリの存在チェック ＆　ディレクトリを表しているかどうか
	if m, err := os.Stat(d.dir); os.IsNotExist(err) {
		return err
	} else if !m.IsDir() {
		return fmt.Errorf("%s : not a directory", d.dir)
	}

	return nil

}

//setFileName urlからファイル名を切り出します、ファイル名を設定します
func (d *downloader) setFileName() {
	token := strings.Split(d.targetURL, "/")

	var fileName string
	for i := 1; fileName == ""; i++ {
		fileName = token[len(token)-i]
	}

	d.fileName = fileName
}

//setFileSize HTTPのヘッダーからファイルサイズを取得し、分割サイズを設定します
func (d *downloader) setFileSize() error {
	resp, err := http.Head(d.targetURL)

	if err != nil {
		return err
	}

	d.fileSize = uint(resp.ContentLength)
	d.split = d.fileSize / d.div

	return nil

}
