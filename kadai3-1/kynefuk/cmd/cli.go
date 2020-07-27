package cmd

import (
	"fmt"
	"io"
	"os"
	"time"
)

// ExitCodeOK represents status ok
// ExitCodeErr represents status error
const (
	ExitCodeOK = iota
	ExitCodeErr
)

// WordsFilePath is a file path of words.txt
// ファイルパスはmain.goからの相対パスじゃないとあかん
const WordsFilePath = "testdata/words.txt"

// CLI is a cli object
type CLI struct {
	OutSteream, ErrStream io.Writer
}

// Run invoke main logic
func (cli *CLI) Run(limit int) int {
	f, err := readWordsFile()
	if err != nil {
		fmt.Fprintf(cli.ErrStream, "failed to read words file. error: %s", err)
		return ExitCodeErr
	}
	questioner := NewQuestioner(f)
	questioner.readSource()
	// select構文を用いて、制限時間と入力を同時に待つ
	c1 := make(chan bool, 1)
	startLimit(limit, c1)
	for {
		select {
		case <-c1:
			break
		}
	}

	return ExitCodeOK
}

func startLimit(limit int, c1 chan bool) {
	now := time.Now()
	timeLimit := now.Add(time.Duration(limit) * time.Second)
	for {
		now = time.Now()
		if timeLimit.After(now) {
			c1 <- true
		}
	}
}

func readWordsFile() (io.Reader, error) {
	f, err := os.Open(WordsFilePath)
	return f, err
}

// NewCLI is a constructor of Questioner
func NewCLI(outStream, errStream io.Writer) *CLI {
	return &CLI{OutSteream: outStream, ErrStream: errStream}
}
