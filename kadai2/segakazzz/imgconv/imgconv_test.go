package imgconv

import (
	"testing"
)

type pattern struct {
	dirname  string
	input    string
	output   string
	expected *converter
	isError  bool
}

func TestNewConverter(t *testing.T) {
	successPatterns := []pattern{
		{"testdata", "png", "jpg", &converter{dirname: "testdata", input: "png", output: "jpg"}, false},
		{"testdata", "jpg", "png", &converter{dirname: "testdata", input: "jpg", output: "png"}, false},
	}
	errorPatterns := []pattern{
		{"testdata", "jpg", "gif", &converter{}, true},
		{"testdata", "gif", "jpg", nil, true},
		{"testdata", "jpg", "jpg", nil, true},
	}

	t.Run("successPatterns", func(t *testing.T) {
		for i, p := range successPatterns {
			testNewConverter(t, i, p)
		}
	})

	t.Run("errorPatterns", func(t *testing.T) {
		for i, p := range errorPatterns {
			testNewConverter(t, i, p)
		}
	})
}

func testNewConverter(t *testing.T, i int, p pattern) {
	t.Helper()
	actual, err := newConverter(p.dirname, p.input, p.output)
	if err != nil && p.isError == false {
		t.Fatalf("pattern[%d]: %v want: NO ERROR, actual: %v", i, p, err)
	}
	if err == nil && p.isError == true {
		t.Fatalf("pattern[%d]: %v want: ERROR, actual: NO ERROR", i, p)
	}
	if (actual == nil && p.expected != nil) || (actual != nil && p.expected == nil) {
		t.Fatalf("pattern[%d]: %v want: isNil(%t), actual: isNil(%t)", i, p, p.expected == nil, actual == nil)
	}

	if actual != nil && p.expected != nil && *actual != *p.expected {
		t.Fatalf("pattern[%d]: %v want: %v [%T], actual: %v [%T]", i, p, *p.expected, *p.expected, *actual, *actual)
	}

}
