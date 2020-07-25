package main

import (
	"bytes"
	"path/filepath"
	"reflect"
	"testing"
)

func TestInput(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{
			name: "match input channel",
			in:   "hoge",
			want: "hoge",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			stdin := bytes.NewBufferString(test.in)

			inc, _ := input(stdin)

			select {
			case in := <-inc:
				if in != test.want {
					t.Errorf(`want="%v" actual="%v"`, test.want, in)
				}
			}
		})
	}
}

func TestImportWords(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		want     []string
	}{
		{
			name:     "import testfile1",
			filepath: filepath.Join("testdata", "words1.txt"),
			want:     []string{"gopher", "golang", "goroutines"},
		},
		{
			name:     "import testfile2",
			filepath: filepath.Join("testdata", "words2.txt"),
			want:     []string{"java", "php", "erlang", "elixir"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, _ := importWords(test.filepath)

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(`filepath="%v" want="%v" actual="%v"`, test.filepath, test.want, got)
			}
		})
	}
}
