package imgconv

import (
	"bytes"
	"image"
	"io"
	"os"
	"testing"
)

type mockEncoder struct{}

func (e *mockEncoder) execute(w io.Writer, Image image.Image) error {
	return nil
}

func TestConverter_Execute(t *testing.T) {
	file, err := os.Open("../testdata/test2/gopher.jpg")
	if err != nil {
		t.Fatal(err)
	}
	stdout := new(bytes.Buffer)

	converter := NewConverter(&mockEncoder{})

	err = converter.Execute(file, stdout)
	if err != nil {
		t.Errorf("failed to call Execute(): %s", err)
	}
}
