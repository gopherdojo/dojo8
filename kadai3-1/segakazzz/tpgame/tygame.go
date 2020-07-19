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
	"sync"
	"time"
)

type game struct {
	in         io.Reader
	out        io.Writer
	buffReader *bufio.Reader
	source     []string
	questions  []singleGame
	matchRatio float64
	timeout    time.Duration
	gch        *gameCh
}

type singleGame struct {
	input           string
	answer          string
	match           bool
	cumulativeMatch int
}

type gameCh struct {
	ch     chan singleGame
	closed bool
	mutex  sync.Mutex
}

func (gc *gameCh) safeClose() {
	gc.mutex.Lock()
	defer gc.mutex.Unlock()
	if !gc.closed {
		close(gc.ch)
		gc.closed = true
	}
}

func (gc *gameCh) isClosed() bool {
	gc.mutex.Lock()
	defer gc.mutex.Unlock()
	return gc.closed
}

func Run(filename string, timeout time.Duration) error {
	g, err := newGame(filename, timeout)
	if err != nil {
		return err
	}

	go func() {
		var sg singleGame
		for i, v := range g.source {
			fmt.Fprintf(g.out, "[%d] %s >>> ", i + 1, v)
			input, _ := g.buffReader.ReadString('\n')
			input = strings.TrimSuffix(input, "\n")

			sg.input = input
			sg.answer = v
			sg.match = (input == v)
			if input == v {
				sg.cumulativeMatch++
			}

			if !g.gch.isClosed() {
				g.gch.ch <- sg
			}
		}
	}()

	select {
	case <-time.After(g.timeout):
		g.displayResult()
		return nil
	}

}

func (g *game) calcResult() {
	nMatch := g.questions[len(g.questions)-1].cumulativeMatch
	nTotal := len(g.questions)
	g.matchRatio = (float64(nMatch) / float64(nTotal)) * 100
}

func (g *game) displayResult() {
	fmt.Fprintln(g.out)
	fmt.Fprintln(g.out, strings.Repeat("-", 80))
	fmt.Fprintln(g.out, "Timeout!")
	g.gch.safeClose()
	for rec := range g.gch.ch {
		g.questions = append(g.questions, rec)
	}
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
	g.calcResult()

	fmt.Fprintf(g.out, "[Summary]\n")
	fmt.Fprintf(g.out, "%-20s %-20s %-20s\n", "Num of Questions", "Num of Correct Ans", "Match Ratio")
	fmt.Fprintln(g.out, strings.Repeat("=", 80))
	fmt.Fprintf(g.out, "%-20d %-20d %.2f\n", len(g.questions), g.questions[len(g.questions)-1].cumulativeMatch, g.matchRatio)
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

	gch := &gameCh{ch: make(chan singleGame, len(words))}
	buffReader := bufio.NewReader(os.Stdin)
	return &game{source: words, timeout: timeout, gch: gch, in: os.Stdin, out: os.Stdout, buffReader: buffReader}, nil
}

func (g *game) display(iw int, to io.Writer) error {
	return nil
}

func (g *game) check(iw int) error {
	// TODO : increment nRightAns, nQuestions

	return nil
}

func (g *game) wait() error {
	return nil
}
