package imgconv_test

import (
	"testing"

	"github.com/gopherdojo/dojo8/kadai2/tanaka0325/imgconv"
)

func TestImgconvNewImageFormat(t *testing.T) {
	tests := []struct {
		args   string
		expect string
	}{
		{args: "png", expect: "png"},
		{args: "jpg", expect: "jpeg"},
		{args: "jpeg", expect: "jpeg"},
		{args: "gif", expect: "gif"},
		{args: "bmp", expect: "bmp"},
		{args: "tiff", expect: "tiff"},
		{args: "tif", expect: "tiff"},
	}

	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			f := imgconv.NewImageFormat(tt.args)
			got := f.GetExt()
			if got != tt.expect {
				t.Errorf("expect = %+v, but got = %+v", tt.expect, got)
			}
		})
	}

	t.Run("unexpected args", func(t *testing.T) {
		got := imgconv.NewImageFormat("pdf")
		if got != nil {
			t.Errorf("expect = nil, but got = %+v", got)
		}
	})
}
