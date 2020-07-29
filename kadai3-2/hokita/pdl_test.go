package pdl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {
	tests := map[string]struct {
		proc int
		url  string
		dir  string
		want string
	}{
		"download": {
			proc: 3,
			url:  "https://blog.golang.org/gopher/header.jpg",
			dir:  "testdata",
			want: filepath.Join("testdata", "header.jpg"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			pdl := New(test.proc, test.url, test.dir)
			err := pdl.Run()

			if err != nil {
				t.Fatal(err)
			}

			if _, err := os.Stat(test.want); err != nil {
				t.Errorf(`"%v" was not found`, test.want)
			}

			if err := os.Remove(test.want); err != nil {
				t.Fatal(err)
			}
		})
	}
}
