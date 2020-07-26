package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

type word string
type words []word

var errOut io.Writer = os.Stderr
var out io.Writer = os.Stdout

//オプション指定
var t = flag.Int("t", 30, "time limit(unit:seconds)")

type result struct {
	word  word
	input string
	match bool
}

func main() {
	flag.Parse()

	//出力する英単語リストの取得
	words, err := getWords()
	if err != nil {
		fmt.Fprintf(errOut, err.Error())
		os.Exit(1)
	}

	//結果を記録するための構造体
	rs := make([]result, 0, 100)
	//結果をやり取りするためのチャネル
	ch := make(chan result)
	//時間の設定
	tm := time.After(time.Duration(*t) * time.Second)

	//MEMO:引数のresultは、スライスのためアドレスを渡すことはしない
	go do(&words, ch)

LOOP:
	for {
		select {
		case r := <-ch:
			rs = append(rs, r)
		case <-tm:
			//結果を表示
			printResult(rs)
			//close(abort)
			break LOOP
		}
	}

}

//getWords 英単語のリストを取得します
func getWords() (words, error) {
	wordsFile := "./words.txt"

	words := make(words, 0, 2000)

	f, err := os.Open(wordsFile)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		words = append(words, word(line))
	}

	if scanner.Err() != nil {
		return nil, err
	}

	return words, nil
}

//do 英単語の表示、入力結果の送信を行います
func do(words *words, ch chan<- result) {
	for {
		//標準出力にランダムで1単語出力する
		word := printWord(words)

		input, err := getInput()
		if err != nil {
			fmt.Fprintf(errOut, err.Error())
			os.Exit(1)
		}

		ch <- result{
			word:  word,
			input: input,
			match: string(word) == input,
		}
	}

}

//printWord ランダムで英単語を出力します
func printWord(words *words) word {
	//ランダムで出力するために、乱数を作成
	t := time.Now().UnixNano()
	rand.Seed(t)

	//indexを指定して取得するため、len(words)までを上限とする
	s := rand.Intn(len(*words))

	fmt.Fprintln(out, (*words)[s])
	return (*words)[s]
}

//getInput 標準入力から１行読み込みます
func getInput() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)

	//読み込み失敗を5回繰り返したら、エラーを返す
	var input string

	for errCnt := 0; errCnt < 5; {
		scanner.Scan()
		input = scanner.Text()

		if err := scanner.Err(); err != nil {
			errCnt++
			fmt.Fprintf(errOut, "読み込みに失敗しました(%d回目): %v", errCnt+1, err)
		} else {
			//問題なく読み込めた場合、読み込み結果を返す
			return input, nil
		}
	}
	return input, fmt.Errorf("読み込みに5回連続で失敗しました")
}

//printResult 結果の表示を行います
func printResult(rs []result) {
	var matchCnt int

	//入力途中でタイムアウトした場合を考慮し、改行を追加しておく
	fmt.Fprintln(out)
	fmt.Fprintf(out, "%-20s %-20s (%s)\n", "表示", "入力", "一致")
	fmt.Fprintln(out, strings.Repeat("-", 75))

	for _, r := range rs {
		var m string

		switch {
		case r.match:
			m = "◯"
			matchCnt++
		case !r.match:
			m = "×"
		}

		fmt.Fprintf(out, "%-20s %-20s (%s)\n", r.word, r.input, m)
	}

	matchRate := float64(matchCnt) / float64(len(rs)) * 100

	fmt.Fprintln(out, strings.Repeat("-", 75))
	fmt.Fprintf(out, "%d問中 %d問 (%.2f %%)\n", len(rs), matchCnt, matchRate)
	fmt.Fprintln(out, strings.Repeat("-", 75))
}
