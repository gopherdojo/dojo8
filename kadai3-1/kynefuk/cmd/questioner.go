package cmd

import (
	"bufio"
	"fmt"
	"io"
)

// Questioner has a role of giving a typing word
type Questioner struct {
	Source io.Reader
}

func (q *Questioner) readSource() {
	scanner := bufio.NewScanner(q.Source)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

// NewQuestioner is a constructor of Questioner
func NewQuestioner(source io.Reader) *Questioner {
	return &Questioner{Source: source}
}
