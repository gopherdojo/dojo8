package main

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestTyping_newTyping(t *testing.T) {
	tests := []struct {
		name      string
		words     []string
		wantWords []string
	}{
		{
			name:      "initialize",
			words:     []string{"a", "b"},
			wantWords: []string{"a", "b"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			typing := newTyping(test.words)

			if typing.word == "" {
				t.Fatal("did not have shuffled word")
			}

			if !reflect.DeepEqual(typing.words, test.wantWords) {
				t.Errorf(`want="%v" actual="%v"`, typing.words, test.wantWords)
			}
		})
	}
}

func TestTyping_check(t *testing.T) {
	tests := []struct {
		name      string
		word      string
		typedWord string
		want      bool
		wantPoint int
	}{
		{
			name:      "match",
			word:      "hoge",
			typedWord: "hoge",
			want:      true,
			wantPoint: 1,
		},
		{
			name:      "unmatch",
			word:      "hoge",
			typedWord: "fuga",
			want:      false,
			wantPoint: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			typing := &Typing{
				word: test.word,
			}

			got := typing.check(test.typedWord)
			if got != test.want {
				t.Errorf(
					`word="%v" typedWord="%v" want="%v" got="%v"`,
					test.word, test.typedWord, test.want, got,
				)
			}

			if typing.point != test.wantPoint {
				t.Errorf(
					`word="%v" typedWord="%v" wantPoint="%v" acutal="%v"`,
					test.word, test.typedWord, test.wantPoint, typing.point,
				)
			}
		})
	}
}

func TestTyping_shuffle(t *testing.T) {
	tests := []struct {
		name string
		seed int64
		want string
	}{
		{
			name: "randseed1",
			seed: 1,
			want: "c",
		},
		{
			name: "randseed2",
			seed: 2,
			want: "b",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			rand.Seed(test.seed)

			typing := &Typing{
				words: []string{"a", "b", "c"},
			}

			typing.shuffle()

			if typing.word != test.want {
				t.Errorf(
					`seed="%v" want="%v" actual="%v"`,
					test.seed, test.want, typing.word,
				)
			}
		})
	}
}
