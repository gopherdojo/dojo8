package imgconv

import (
	"reflect"
	"testing"
)

func TestNewImage(t *testing.T) {
	tests := map[string]struct {
		ext        string
		want       Image
		wantErrMsg string
	}{
		"jpg": {
			ext:  ".jpg",
			want: &JpegImage{},
		},
		"png": {
			ext:  ".png",
			want: &PngImage{},
		},
		"txt": {
			ext:        ".txt",
			wantErrMsg: "selected extension is not supported",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := NewImage(test.ext)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf(
					`ext="%v" want="%v" got="%v"`,
					test.ext, test.want, got,
				)
			}

			if err != nil && err.Error() != test.wantErrMsg {
				t.Errorf(
					`ext="%v" wantErrMsg="%v" errMsg="%v"`,
					test.ext, test.wantErrMsg, err.Error(),
				)
			}
		})
	}
}
