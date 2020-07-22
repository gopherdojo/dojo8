package imgconv_test

import (
	"fmt"
	"image"
	"io"
	"os"
	"testing"

	"github.com/gopherdojo/dojo8/kadai2/tanaka0325/imgconv"
)

func TestImgConvDo(t *testing.T) {
	tests := []struct {
		name  string
		args  imgconv.ConvertParam
		isErr bool
	}{
		{
			name: "error: Open failed",
			args: func() imgconv.ConvertParam {
				mf := newMockFileHandler(
					func(s string) (io.ReadCloser, error) { return nil, fmt.Errorf("err") },
					nil,
				)

				return imgconv.ConvertParam{
					Path:        "path",
					FileHandler: mf,
				}
			}(),
			isErr: true,
		},
		{
			name: "error: Create failed",
			args: func() imgconv.ConvertParam {
				mf := newMockFileHandler(
					func(s string) (io.ReadCloser, error) { return mockCloser{}, nil },
					func(s string) (io.WriteCloser, error) { return nil, fmt.Errorf("err") },
				)

				mbi := newMockImageFormat(
					func(r io.Reader) (image.Image, error) { return mockImage{}, nil },
					nil,
					func() string { return "jpeg" },
				)

				mai := newMockImageFormat(
					nil,
					func(w io.Writer, i image.Image) error { return nil },
					func() string { return "png" },
				)

				return imgconv.ConvertParam{
					Path:         "path",
					FileHandler:  mf,
					BeforeFormat: mbi,
					AfterFormat:  mai,
				}
			}(),
			isErr: true,
		},
		{
			name: "error: convert failed",
			args: func() imgconv.ConvertParam {
				mf := newMockFileHandler(
					func(s string) (io.ReadCloser, error) { return mockCloser{}, nil },
					func(s string) (io.WriteCloser, error) { return mockCloser{}, nil },
				)

				mbi := newMockImageFormat(
					func(r io.Reader) (image.Image, error) { return mockImage{}, nil },
					nil,
					func() string { return "jpeg" },
				)

				mai := newMockImageFormat(
					nil,
					func(w io.Writer, i image.Image) error { return fmt.Errorf("err") },
					func() string { return "png" },
				)

				return imgconv.ConvertParam{
					Path:         "path",
					FileHandler:  mf,
					BeforeFormat: mbi,
					AfterFormat:  mai,
				}
			}(),
			isErr: true,
		},
		{
			name: "ok",
			args: func() imgconv.ConvertParam {
				mf := newMockFileHandler(
					func(s string) (io.ReadCloser, error) { return mockCloser{}, nil },
					func(s string) (io.WriteCloser, error) { return mockCloser{}, nil },
				)

				mbi := newMockImageFormat(
					func(r io.Reader) (image.Image, error) { return mockImage{}, nil },
					nil,
					func() string { return "jpeg" },
				)
				mai := newMockImageFormat(
					nil,
					func(w io.Writer, i image.Image) error { return nil },
					func() string { return "png" },
				)

				return imgconv.ConvertParam{
					Path:         "path",
					FileHandler:  mf,
					BeforeFormat: mbi,
					AfterFormat:  mai,
				}
			}(),
			isErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := imgconv.Do(tt.args)
			if (tt.isErr && err == nil) || (!tt.isErr && err != nil) {
				t.Errorf("expect err = %+v, but got err = %+v", tt.isErr, err)
			}
		})
	}
}

func TestImgConv_convert(t *testing.T) {
	type args struct {
		Reader  io.Reader
		Decoder imgconv.Decoder
		Writer  io.Writer
		Encoder imgconv.Encoder
	}

	tests := []struct {
		name  string
		args  args
		isErr bool
	}{
		{
			name: "error: Decode failed",
			args: func() args {
				mi := newMockImageFormat(
					func(r io.Reader) (image.Image, error) { return nil, fmt.Errorf("err") },
					nil,
					nil,
				)
				return args{
					os.Stdin,
					mi,
					nil,
					nil,
				}
			}(),
			isErr: true,
		},
		{
			name: "error: Encode failed",
			args: func() args {
				me := newMockImageFormat(
					func(r io.Reader) (image.Image, error) { return mockImage{}, nil },
					nil,
					nil,
				)

				md := newMockImageFormat(
					nil,
					func(w io.Writer, i image.Image) error { return fmt.Errorf("err") },
					nil,
				)
				return args{
					os.Stdin,
					me,
					os.Stdout,
					md,
				}
			}(),
			isErr: true,
		},
		{
			name: "ok",
			args: func() args {
				me := newMockImageFormat(
					func(r io.Reader) (image.Image, error) { return mockImage{}, nil },
					nil,
					nil,
				)

				md := newMockImageFormat(
					nil,
					func(w io.Writer, i image.Image) error { return nil },
					nil,
				)

				return args{
					os.Stdin,
					me,
					os.Stdout,
					md,
				}
			}(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := imgconv.ConvertFunc(tt.args.Reader, tt.args.Decoder, tt.args.Writer, tt.args.Encoder)
			if (tt.isErr && err == nil) || (!tt.isErr && err != nil) {
				t.Errorf("expect err = %+v, but got err = %+v", tt.isErr, err)
			}
		})
	}
}
