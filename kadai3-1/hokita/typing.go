package main

import (
	"math/rand"
)

type Typing struct {
	point int
	word  string
	words []string
}

func newTyping(words []string) *Typing {
	typing := &Typing{
		words: words,
	}

	typing.shuffle()

	return typing
}

func (t *Typing) pointUp() {
	t.point += 1
}

func (t *Typing) getPoint() int {
	return t.point
}

func (t *Typing) check(word string) bool {
	if t.word == word {
		t.pointUp()
		return true
	}
	return false
}

func (t *Typing) shuffle() {
	i := rand.Intn(len(t.words))
	t.word = t.words[i]
}
