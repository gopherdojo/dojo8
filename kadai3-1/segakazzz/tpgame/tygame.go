package tpgame

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type game struct {
	in         io.Reader
	out        io.Writer
	buffReader *bufio.Reader
	source     []string
	questions  []singleGame
	matchRatio float64
	numMatch   int
	timeout    time.Duration
	errCh      chan error
}

type singleGame struct {
	input  string
	answer string
	match  bool
}

func Run(filename string, timeout time.Duration) error {
	g, err := newGame(filename, timeout)
	if err != nil {
		return err
	}
	go func() {
		err := g.processGame()
		if err !=nil {
			g.errCh <- err
		}
	}()

	select {
	case err=<-g.errCh:
		return err
	case <-time.After(g.timeout):
		g.displayResult()
		return nil
	}

}

func (g *game) processGame() error{

	var sg singleGame
	for i, v := range g.source {
		fmt.Fprintf(g.out, "[%d] %s >>> ", i+1, v)
		input, err := g.buffReader.ReadString('\n')
		if err != nil {
			return err
		}
		input = strings.TrimSuffix(input, "\n")
		sg.input = input
		sg.answer = v
		sg.match = (input == v)
		g.updateSingleGameResult(sg)
	}

	return nil

}
func (g *game) updateSingleGameResult(sg singleGame) {
	g.questions = append(g.questions, sg)
	if sg.match {
		g.numMatch++
	}
	g.matchRatio = (float64(g.numMatch) / float64(len(g.questions))) * 100
}

func (g *game) displayResult() {
	fmt.Fprintln(g.out)
	fmt.Fprintln(g.out, strings.Repeat("-", 80))
	fmt.Fprintln(g.out, "Timeout!")
	fmt.Fprintf(g.out, "%-10s %-20s %-20s %-5s\n", "#", "Your Input", "Answer", "Correct?")
	fmt.Fprintln(g.out, strings.Repeat("=", 80))
	for i, q := range g.questions {
		var correct rune
		if q.match {
			correct = '⭕'
		} else {
			correct = '❌'
		}
		fmt.Fprintf(g.out, "%-10d %-20s %-20s %-5c\n", i+1, q.input, q.answer, correct)
	}

	fmt.Fprintf(g.out, "[Summary]\n")
	fmt.Fprintf(g.out, "%-20s %-20s %-20s %-20s\n", "Num of Questions", "Num of Correct ANS", "Match Ratio[%]", "Timeout Duration[sec]")
	fmt.Fprintln(g.out, strings.Repeat("=", 80))
	fmt.Fprintf(g.out, "%-20d %-20d %-20.2f %-20s\n", len(g.questions), g.numMatch, g.matchRatio, g.timeout)
}

func newGame(filename string, timeout time.Duration) (g *game, err error) {
	var words []string
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = jsonFile.Close()
	}()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(byteValue, &words)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(words), func(i, j int) { words[i], words[j] = words[j], words[i] })

	buffReader := bufio.NewReader(os.Stdin)

	errCh := make(chan error)
	return &game{
		source: words,
		timeout: timeout,
		in: os.Stdin,
		out: os.Stdout,
		buffReader: buffReader,
		errCh:errCh,
	}, nil
}
