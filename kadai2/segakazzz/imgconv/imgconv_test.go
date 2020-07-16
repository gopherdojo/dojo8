package imgconv

import (
	"testing"
)

// func isSameConverter(actual *converter, expected *converter) bool {
// 	if actual == nil && expected == nil {
// 		return true
// 	}
// 	if actual != nil && expected == nil {
// 		return false
// 	}
// 	if actual == nil && expected != nil {
// 		return false
// 	}
// 	if actual.dirname != expected.dirname {
// 		return false
// 	}
// 	if actual.input != expected.input {
// 		return false
// 	}
// 	if actual.output != expected.output {
// 		return false
// 	}
// 	return true
// }

func TestNewConverter(t *testing.T) {
	patterns := []struct {
		dirname  string
		input    string
		output   string
		expected *converter
		isError  bool
	}{
		{"testdata", "png", "jpg", &converter{dirname: "testdata", input: "png", output: "jpg"}, false},
		{"testdata", "jpg", "png", &converter{dirname: "testdata", input: "jpg", output: "png"}, false},
		{"testdata", "jpg", "gif", nil, true},
		{"testdata", "gif", "jpg", nil, true},
		{"testdata", "jpg", "jpg", nil, true},
	}

	for i, p := range patterns {
		actual, err := newConverter(p.dirname, p.input, p.output)
		if err != nil && p.isError == false {
			t.Errorf("pattern: %d want: NO ERROR, actual: %v", i, err)
		}
		if err == nil && p.isError == true {
			t.Errorf("pattern: %d want: ERROR, actual: NO ERROR", i)
		}

		if actual != nil && p.expected != nil && *actual != *p.expected {
			t.Errorf("pattern: %d want: %v [%T], actual: %v [%T]", i, *actual, *actual, *p.expected, *p.expected)
		}
	}
}
