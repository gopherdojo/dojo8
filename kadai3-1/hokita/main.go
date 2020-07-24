package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

const (
	ExitCodeOK    = 0
	ExitCodeError = 1

	timelimit = 10
	filePath  = "./words.txt"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	exitCode := run()
	os.Exit(exitCode)
}

func run() int {
	in, errc := input(os.Stdin)
	timelimit := time.After(timelimit * time.Second)

	words, err := importWords(filePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return ExitCodeError
	}

	typing := newTyping(words)

	for {
		fmt.Println(typing.word)
		fmt.Print(">")

		select {
		case word := <-in:
			if typing.check(word) {
				fmt.Println("Good!")
				fmt.Println()
			} else {
				fmt.Println("Bad..")
				fmt.Println()
			}
			typing.shuffle()
		case err := <-errc:
			fmt.Println()
			fmt.Fprintln(os.Stderr, err)
			return ExitCodeError
		case <-timelimit:
			fmt.Println()
			fmt.Println("------")
			fmt.Println("Finish!!")
			fmt.Printf("Result: %d points\n", typing.getPoint())
			return ExitCodeOK
		}
	}
}

func input(r io.Reader) (<-chan string, <-chan error) {
	// チャネルを作る
	result := make(chan string)
	errc := make(chan error)

	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			// チャネルに読み込んだ文字列を送る
			result <- s.Text()
			if err := s.Err(); err != nil {
				errc <- err
			}
		}
		// チャネルを閉じる
		close(result)
	}()
	// チャネルを返す
	return result, errc
}

// cf. https://qiita.com/jpshadowapps/items/ae7274ec0d40882d76b5
func importWords(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines := make([]string, 0, 10)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if serr := scanner.Err(); serr != nil {
		return nil, err
	}

	return lines, nil
}
