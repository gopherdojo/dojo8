package imgconv

import (
	"errors"
	"testing"
)

func TestIsValidFileType(t *testing.T) {
	t.Parallel()
	testTable := []struct {
		desc   string
		arg    string
		expect bool
	}{
		{desc: "jpegがtrueで返される", arg: "jpeg", expect: true},
		{desc: "pngがtrueで返される", arg: "png", expect: true},
		{desc: "jpgがfalseで返される", arg: "jpg", expect: false},
		{desc: "gifがfalseで返される", arg: "gif", expect: false},
		{desc: "空文字がfalseで返される", arg: "", expect: false},
	}

	for _, test := range testTable {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			actual := IsValidFileType(test.arg)
			if test.expect != actual {
				t.Errorf("%s 失敗\n期待: %v\n実際: %v\n", t.Name(), test.expect, actual)
			}
		})
	}
}

func TestIsDir(t *testing.T) {
	t.Parallel()
	testTable := []struct {
		desc    string
		arg     string
		expect1 bool
		expect2 error
	}{
		{desc: "ディレクトリを指定したらtrue", arg: "../testdata", expect1: true},
		{desc: "ファイルを指定したらfalse", arg: "../testdata/Example.jpg", expect1: false},
		{desc: "存在しないパスを指定したらエラー", arg: "./testdata/NotFound.jpg", expect2: FileStatError},
		{desc: "空文字を指定したらエラー", arg: "", expect2: FileStatError},
		// {desc: "不良セクタを指定したらfalse", arg: "", expect1: false}, // やり方わからない
	}

	for _, test := range testTable {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			actual1, actual2 := IsDir(test.arg)
			if test.expect1 != actual1 || !errors.Is(actual2, test.expect2) {
				t.Errorf("%s 失敗\n期待: %v, %v\n実際: %v, %v\n", t.Name(), test.expect1, test.expect2, actual1, actual2)
			}
		})
	}
}
