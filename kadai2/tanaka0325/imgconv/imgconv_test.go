package imgconv_test

import (
	"image"
	"io"
	"testing"

	"github.com/gopherdojo/dojo8/kadai2/tanaka0325/imgconv"
)

type mockDecoder struct{}
type mockEncoder struct{}

func (m mockDecoder) Decode(io.Reader) (image.Image, error) { return nil, nil }
func (m mockEncoder) Encode(io.Writer, image.Image) error   { return nil }

func TestImgConvDo(t *testing.T) {
	md := mockDecoder{}
	me := mockEncoder{}

	tests := []struct {
		name  string
		args  imgconv.ConvertParam
		isErr bool
	}{
		{
			name: "ok",
			args: imgconv.ConvertParam{
				BeforeImage: md,
				AfterImage:  me,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := imgconv.Do(tt.args)
			if (tt.isErr && err == nil) || (!tt.isErr && err != nil) {
				t.Errorf("expect err = %t, but got err = %s", tt.isErr, err)
			}
		})
	}
}
