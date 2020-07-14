package imageconverter

import (
	"errors"
	"fmt"
	"image"
	"io"
	"os"
	"reflect"
	"testing"
	"time"
)

func Test_replaceExt(t *testing.T) {
	type args struct {
		filePath string
		toExt    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Normal", args{"hoge/fuga.jpg", ".png"}, "hoge/fuga.png"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceExt(tt.args.filePath, tt.args.toExt); got != tt.want {
				t.Errorf("replaceExt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Convert(t *testing.T) {
	tests := []struct {
		name      string
		args      Args
		convImage func(path string, info os.FileInfo, err error, args Args) error
		wantErr   bool
	}{
		{
			name: "Normal",
			args: Args{"./hoge", "jpg", "png"},
			convImage: func(path string, info os.FileInfo, err error, args Args) error {
				return nil
			},
			wantErr: false,
		},
		{
			name: "ErrorBecauseWalkReturnError",
			args: Args{"./hoge", "jpg", "png"},
			convImage: func(path string, info os.FileInfo, err error, args Args) error {
				return fmt.Errorf("error")
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			convImage = tt.convImage

			if err := Convert(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("Convert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_convertImage(t *testing.T) {
	type args struct {
		path string
		info os.FileInfo
		err  error
		args Args
	}
	tests := []struct {
		name    string
		args    args
		e       func(source, dest string) (err error)
		wantErr bool
	}{
		{
			name: "Normal",
			args: args{
				path: "./hoge/fuga.jpg",
				info: FileInfoMock{},
				err:  nil,
				args: Args{
					Directory:       "./hoge",
					BeforeExtension: "jpg",
					AfterExtension:  "png",
				},
			},
			e: func(source, dest string) (err error) {
				return nil
			},
			wantErr: false,
		},
		{
			name: "NormalTargetIsDir",
			args: args{
				path: "./hoge/fuga.jpg",
				info: FileInfoIsDirTrueMock{},
				err:  nil,
				args: Args{
					Directory:       "./hoge",
					BeforeExtension: "jpg",
					AfterExtension:  "png",
				},
			},
			e: func(source, dest string) (err error) {
				return nil
			},
			wantErr: false,
		},
		{
			name: "NormalFileNameIsNotMatch",
			args: args{
				path: "./hoge/fuga.jpg",
				info: FileInfoMock{},
				err:  nil,
				args: Args{
					Directory:       "./hoge",
					BeforeExtension: "gif",
					AfterExtension:  "png",
				},
			},
			e: func(source, dest string) (err error) {
				return nil
			},
			wantErr: false,
		},
		{
			name: "ErrorBecauseDirectoryError",
			args: args{
				path: "./hoge/fuga.jpg",
				info: FileInfoMock{},
				err:  errors.New("error"),
				args: Args{
					Directory:       "./hoge",
					BeforeExtension: "jpg",
					AfterExtension:  "png",
				},
			},
			e: func(source, dest string) (err error) {
				return nil
			},
			wantErr: true,
		},
		{
			name: "ErrorBecauseExecReturnError",
			args: args{
				path: "./hoge/fuga.jpg",
				info: FileInfoMock{},
				err:  nil,
				args: Args{
					Directory:       "./hoge",
					BeforeExtension: "jpg",
					AfterExtension:  "png",
				},
			},
			e: func(source, dest string) (err error) {
				return fmt.Errorf("error")
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			execute = tt.e
			if err := convertImage(tt.args.path, tt.args.info, tt.args.err, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("convertImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_exec(t *testing.T) {
	type mockFns struct {
		osOpen          func(name string) (*os.File, error)
		osCreate        func(name string) (*os.File, error)
		sourceFileClose func(file *os.File) error
		destFileClose   func(file *os.File) error
		imageDecode     func(r io.Reader) (image.Image, string, error)
		newImgEncoder   func(ext string) (Encoder, error)
	}
	type args struct {
		source string
		dest   string
	}
	tests := []struct {
		name    string
		args    args
		mockFns mockFns
		wantErr bool
	}{
		{
			name: "Normal",
			args: args{source: "jpeg", dest: "png"},
			mockFns: mockFns{
				osOpen: func(name string) (*os.File, error) {
					return nil, nil
				},
				sourceFileClose: func(file *os.File) error {
					return nil
				},
				osCreate: func(name string) (*os.File, error) {
					return nil, nil
				},
				destFileClose: func(file *os.File) error {
					return nil
				},
				imageDecode: func(r io.Reader) (image.Image, string, error) {
					return nil, "", nil
				},
				newImgEncoder: func(ext string) (Encoder, error) {
					return EncoderMock{}, nil
				},
			},
			wantErr: false,
		},
		{
			name: "ErrorBecauseSourceFileOpenReturnError",
			args: args{source: "jpeg", dest: "png"},
			mockFns: mockFns{
				osOpen: func(name string) (*os.File, error) {
					return nil, fmt.Errorf("error")
				},
				sourceFileClose: func(file *os.File) error {
					return nil
				},
				osCreate: func(name string) (*os.File, error) {
					return nil, nil
				},
				destFileClose: func(file *os.File) error {
					return nil
				},
				imageDecode: func(r io.Reader) (image.Image, string, error) {
					return nil, "", nil
				},
				newImgEncoder: func(ext string) (Encoder, error) {
					return EncoderMock{}, nil
				},
			},
			wantErr: true,
		},
		{
			name: "ErrorBecauseDestFileCreateReturnError",
			args: args{source: "jpeg", dest: "png"},
			mockFns: mockFns{
				osOpen: func(name string) (*os.File, error) {
					return nil, nil
				},
				sourceFileClose: func(file *os.File) error {
					return nil
				},
				osCreate: func(name string) (*os.File, error) {
					return nil, fmt.Errorf("error")
				},
				destFileClose: func(file *os.File) error {
					return nil
				},
				imageDecode: func(r io.Reader) (image.Image, string, error) {
					return nil, "", nil
				},
				newImgEncoder: func(ext string) (Encoder, error) {
					return EncoderMock{}, nil
				},
			},
			wantErr: true,
		},
		{
			name: "ErrorBecauseSourceFileDecodeReturnError",
			args: args{source: "jpeg", dest: "png"},
			mockFns: mockFns{
				osOpen: func(name string) (*os.File, error) {
					return nil, nil
				},
				sourceFileClose: func(file *os.File) error {
					return nil
				},
				osCreate: func(name string) (*os.File, error) {
					return nil, nil
				},
				destFileClose: func(file *os.File) error {
					return nil
				},
				imageDecode: func(r io.Reader) (image.Image, string, error) {
					return nil, "", fmt.Errorf("error")
				},
				newImgEncoder: func(ext string) (Encoder, error) {
					return EncoderMock{}, nil
				},
			},
			wantErr: true,
		},
		{
			name: "NormalBecauseNewImageEncoderReturnError",
			args: args{source: "jpeg", dest: "png"},
			mockFns: mockFns{
				osOpen: func(name string) (*os.File, error) {
					return nil, nil
				},
				sourceFileClose: func(file *os.File) error {
					return nil
				},
				osCreate: func(name string) (*os.File, error) {
					return nil, nil
				},
				destFileClose: func(file *os.File) error {
					return nil
				},
				imageDecode: func(r io.Reader) (image.Image, string, error) {
					return nil, "", nil
				},
				newImgEncoder: func(ext string) (Encoder, error) {
					return nil, fmt.Errorf("error")
				},
			},
			wantErr: true,
		},
		{
			name: "NormalBecauseEncodeReturnError",
			args: args{source: "jpeg", dest: "png"},
			mockFns: mockFns{
				osOpen: func(name string) (*os.File, error) {
					return nil, nil
				},
				sourceFileClose: func(file *os.File) error {
					return nil
				},
				osCreate: func(name string) (*os.File, error) {
					return nil, nil
				},
				destFileClose: func(file *os.File) error {
					return nil
				},
				imageDecode: func(r io.Reader) (image.Image, string, error) {
					return nil, "", nil
				},
				newImgEncoder: func(ext string) (Encoder, error) {
					return &EncoderErrorMock{}, nil
				},
			},
			wantErr: true,
		},
		{
			name: "NormalBecauseDeferredSourceFileCloseReturnError",
			args: args{source: "jpeg", dest: "png"},
			mockFns: mockFns{
				osOpen: func(name string) (*os.File, error) {
					return nil, nil
				},
				sourceFileClose: func(file *os.File) error {
					return fmt.Errorf("error")
				},
				osCreate: func(name string) (*os.File, error) {
					return nil, nil
				},
				destFileClose: func(file *os.File) error {
					return nil
				},
				imageDecode: func(r io.Reader) (image.Image, string, error) {
					return nil, "", nil
				},
				newImgEncoder: func(ext string) (Encoder, error) {
					return EncoderMock{}, nil
				},
			},
			wantErr: false,
		},
		{
			name: "ErrorBecauseDeferredDestFileCloseReturnError",
			args: args{source: "jpeg", dest: "png"},
			mockFns: mockFns{
				osOpen: func(name string) (*os.File, error) {
					return nil, nil
				},
				sourceFileClose: func(file *os.File) error {
					return fmt.Errorf("error")
				},
				osCreate: func(name string) (*os.File, error) {
					return nil, nil
				},
				destFileClose: func(file *os.File) error {
					return fmt.Errorf("error")
				},
				imageDecode: func(r io.Reader) (image.Image, string, error) {
					return nil, "", nil
				},
				newImgEncoder: func(ext string) (Encoder, error) {
					return EncoderMock{}, nil
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			osOpen = tt.mockFns.osOpen
			sourceFileClose = tt.mockFns.sourceFileClose
			osCreate = tt.mockFns.osCreate
			destFileClose = tt.mockFns.destFileClose
			imageDecode = tt.mockFns.imageDecode
			newImgEncoder = tt.mockFns.newImgEncoder
			if err := exec(tt.args.source, tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("exec() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newImageEncoder(t *testing.T) {
	type args struct {
		dest string
	}
	tests := []struct {
		name    string
		args    args
		want    Encoder
		wantErr bool
	}{
		{name: "NormalJpeg", args: args{dest: "./hoge/fuga.jpeg"}, want: JpegEncoder{}, wantErr: false},
		{name: "NormalJpg", args: args{dest: "./hoge/fuga.jpg"}, want: JpegEncoder{}, wantErr: false},
		{name: "NormalPng", args: args{dest: "./hoge/fuga.png"}, want: PngEncoder{}, wantErr: false},
		{name: "NormalGif", args: args{dest: "./hoge/fuga.gif"}, want: GifEncoder{}, wantErr: false},
		{name: "NormalTiff", args: args{dest: "./hoge/fuga.tiff"}, want: TiffEncoder{}, wantErr: false},
		{name: "NormalTif", args: args{dest: "./hoge/fuga.tif"}, want: TiffEncoder{}, wantErr: false},
		{name: "NormalBmp", args: args{dest: "./hoge/fuga.bmp"}, want: BmpEncoder{}, wantErr: false},
		{name: "ErrorBecauseInvalidExtension", args: args{dest: "./hoge/fuga.pdf"}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newImageEncoder(tt.args.dest)
			if (err != nil) != tt.wantErr {
				t.Errorf("newImageEncoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newImageEncoder() got = %v, want %v", got, tt.want)
			}
		})
	}
}

type EncoderMock struct{}

func (e EncoderMock) Encode(w io.Writer, m image.Image) error {
	return nil
}

type EncoderErrorMock struct{}

func (e EncoderErrorMock) Encode(w io.Writer, m image.Image) error {
	return fmt.Errorf("error")
}

type FileInfoMock struct{}

func (FileInfoMock) Name() string {
	return "hoge.jpg"
}

func (FileInfoMock) Size() int64 {
	panic("implement me")
}

func (FileInfoMock) Mode() os.FileMode {
	panic("implement me")
}

func (FileInfoMock) ModTime() time.Time {
	panic("implement me")
}

func (FileInfoMock) IsDir() bool {
	return false
}

func (FileInfoMock) Sys() interface{} {
	panic("implement me")
}

type FileInfoIsDirTrueMock struct{}

func (FileInfoIsDirTrueMock) Name() string {
	return "hoge.jpg"
}

func (FileInfoIsDirTrueMock) Size() int64 {
	panic("implement me")
}

func (FileInfoIsDirTrueMock) Mode() os.FileMode {
	panic("implement me")
}

func (FileInfoIsDirTrueMock) ModTime() time.Time {
	panic("implement me")
}

func (FileInfoIsDirTrueMock) IsDir() bool {
	return true
}

func (FileInfoIsDirTrueMock) Sys() interface{} {
	panic("implement me")
}
