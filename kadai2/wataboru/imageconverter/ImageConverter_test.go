package imageconverter

import (
	"errors"
	"os"
	"path/filepath"
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
		{"normal", args{"hoge/fuga.jpg", ".png"}, "hoge/fuga.png"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceExt(tt.args.filePath, tt.args.toExt); got != tt.want {
				t.Errorf("replaceExt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvert(t *testing.T) {
	type args struct {
		args         Args
		FilepathWalk FilepathWalk
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"Normal",
			args{
				Args{"hoge", "jpg", "png"},
				FilepathWalk{func(root string, walkFn filepath.WalkFunc) error {
					return nil
				}},
			},
			false,
		},
		{
			"WalkReturnError",
			args{
				Args{"hoge", "jpg", "png"},
				FilepathWalk{func(root string, walkFn filepath.WalkFunc) error {
					return errors.New("Error")
				}},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Convert(tt.args.args, tt.args.FilepathWalk); (err != nil) != tt.wantErr {
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
		wantErr bool
	}{
		{
			"Normal",
			args{
				"./hoge",
				FileInfoMock{},
				nil,
				Args{
					"./hoge",
					"jpg",
					"png",
				},
			},
			false,
		},
		{
			"Normal",
			args{
				"./hoge",
				FileInfoMock{},
				errors.New("error"),
				Args{
					"./hoge",
					"jpg",
					"png",
				},
			},
			true,
		},
		{
			"DirectoryError",
			args{
				"./hoge",
				FileInfoMock{},
				errors.New("error"),
				Args{
					"./hoge",
					"jpg",
					"png",
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := convertImage(tt.args.path, tt.args.info, tt.args.err, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("convertImage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
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
