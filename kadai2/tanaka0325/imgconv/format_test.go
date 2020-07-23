package imgconv_test

import (
	"testing"

	"github.com/gopherdojo/dojo8/kadai2/tanaka0325/imgconv"
)

func TestImgconvNewImageFormat(t *testing.T) {
	tests := []struct {
		args     string
		expected string
	}{
		{args: "png", expected: "png"},
		{args: "jpg", expected: "jpeg"},
		{args: "jpeg", expected: "jpeg"},
		{args: "gif", expected: "gif"},
		{args: "bmp", expected: "bmp"},
		{args: "tiff", expected: "tiff"},
		{args: "tif", expected: "tiff"},
	}

	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			f := imgconv.NewImageFormat(tt.args)
			got := f.GetExt()
			if got != tt.expected {
				t.Errorf("expected = %+v, but got = %+v", tt.expected, got)
			}
		})
	}

	t.Run("expected args", func(t *testing.T) {
		got := imgconv.NewImageFormat("pdf")
		if got != nil {
			t.Errorf("expected = nil, but got = %+v", got)
		}
	})
}
