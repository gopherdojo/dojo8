package imgconv_test

import (
	"testing"

	"github.com/gopherdojo/dojo8/kadai2/tanaka0325/imgconv"
)

func TestImgconvNewFile(t *testing.T) {
	var f interface{}
	f = imgconv.NewFile()

	got, ok := f.(imgconv.File)
	if !ok {
		t.Errorf("expect type: File, but got %T", got)
	}
}
