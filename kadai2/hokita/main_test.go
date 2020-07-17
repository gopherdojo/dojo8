package main

import (
	"flag"
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	tests := map[string]struct {
		dir  string
		from string
		to   string
		want int
	}{
		"success": {
			dir:  "testdata/test1/",
			from: "jpg",
			to:   "png",
			want: ExitCodeOk,
		},
		"non exit dir": {
			dir:  "testdata/test99/",
			from: "jpg",
			to:   "png",
			want: ExitCodeError,
		},
		"error": {
			from: "jpg",
			to:   "png",
			want: ExitCodeError,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			flag.CommandLine.Set("from", test.from)
			flag.CommandLine.Set("to", test.to)

			if got := run(test.dir); got != test.want {
				t.Errorf(
					`dir="%v" from="%v" to="%v" want="%v" actual="%v"`,
					test.dir, test.from, test.to, test.want, got,
				)
			}

			if test.want != ExitCodeOk {
				return
			}

			outFile := "testdata/test1/gopher.png"
			if err := os.Remove(outFile); err != nil {
				t.Fatal(err)
			}
		})
	}
}
