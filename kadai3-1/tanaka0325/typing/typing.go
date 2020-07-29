package typing

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

type Game struct {
	Words []string
	Time  time.Duration
	Score int
}

func NewGame(ws []string, sec int) *Game {
	t := time.Duration(sec) * time.Second

	return &Game{
		Words: ws,
		Time:  t,
	}
}

func (g *Game) Play() error {
	fmt.Println("game start :)")

	word := choiceWord(g.Words)
	fmt.Printf("word: [%s]\n", word)

	ticker := time.After(g.Time)

	for {
		select {
		case <-ticker:
			fmt.Printf("\ntime up!\n")
			fmt.Printf("Your Score: %d\n", g.Score)
			return nil
		case s := <-input(os.Stdin):
			if word == s {
				fmt.Println("correct!")
				g.Score++
			} else {
				fmt.Println("wrong...")
			}

			word = choiceWord(g.Words)
			fmt.Printf("word: [%s]\n", word)
		}
	}
}

func choiceWord(words []string) string {
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(words))

	return words[i]
}

func input(r io.Reader) <-chan string {
	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(r)
		s.Scan()
		ch <- s.Text()
		close(ch)
	}()
	return ch
}
