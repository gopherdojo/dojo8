package main

import (
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/gopherdojo/dojo8/kadai1/wataboru/imageconverter"
)

type mockFns struct {
	osStat       func(name string) (os.FileInfo, error)
	osIsNotExist func(err error) bool
	imgconv      func(args imageconverter.Args) error
}

func Test_run(t *testing.T) {
	tests := []struct {
		name  string
		args  imageconverter.Args
		mocks mockFns
		want  int
	}{
		{
			name: "Normal",
			args: imageconverter.Args{
				Directory:       "./hoge",
				BeforeExtension: "jpg",
				AfterExtension:  "png",
			},
			mocks: mockFns{
				osStat: func(name string) (os.FileInfo, error) {
					return nil, nil
				},
				osIsNotExist: func(err error) bool {
					return false
				},
				imgconv: func(args imageconverter.Args) error {
					return nil
				},
			},
			want: ExitCodeSuccess,
		},
		{
			name: "ErrorBecauseWithoutArgsDirectory",
			args: imageconverter.Args{
				Directory:       "",
				BeforeExtension: "jpg",
				AfterExtension:  "png",
			},
			mocks: mockFns{
				osStat: func(name string) (os.FileInfo, error) {
					return nil, nil
				},
				osIsNotExist: func(err error) bool {
					return false
				},
				imgconv: func(args imageconverter.Args) error {
					return nil
				},
			},
			want: ExitCodeInvalidDirectoryError,
		},
		{
			name: "ErrorBecauseOsIsNotExitReturnError",
			args: imageconverter.Args{
				Directory:       "./hoge",
				BeforeExtension: "jpg",
				AfterExtension:  "png",
			},
			mocks: mockFns{
				osStat: func(name string) (os.FileInfo, error) {
					return nil, nil
				},
				osIsNotExist: func(err error) bool {
					return true
				},
				imgconv: func(args imageconverter.Args) error {
					return nil
				},
			},
			want: ExitCodeInvalidDirectoryError,
		},
		{
			name: "ErrorBecauseImgconvReturnError",
			args: imageconverter.Args{
				Directory:       "./hoge",
				BeforeExtension: "jpg",
				AfterExtension:  "png",
			},
			mocks: mockFns{
				osStat: func(name string) (os.FileInfo, error) {
					return nil, nil
				},
				osIsNotExist: func(err error) bool {
					return false
				},
				imgconv: func(args imageconverter.Args) error {
					return fmt.Errorf("error")
				},
			},
			want: ExitCodeError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flagsSet(t, tt.args)
			MockSet(t, tt.mocks)
			if got := run(); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}

func flagsSet(t *testing.T, args imageconverter.Args) {
	t.Helper()
	if err := flag.CommandLine.Set("dir", args.Directory); err != nil {
		t.Errorf("error occurred in flagSet helper: %v", err)

	}
	if err := flag.CommandLine.Set("before", args.BeforeExtension); err != nil {
		t.Errorf("error occurred in flagSet helper: %v", err)

	}
	if err := flag.CommandLine.Set("after", args.AfterExtension); err != nil {
		t.Errorf("error occurred in flagSet helper: %v", err)

	}
}

func MockSet(t *testing.T, m mockFns) {
	t.Helper()
	osStat = m.osStat
	osIsNotExist = m.osIsNotExist
	imgconv = m.imgconv
}
