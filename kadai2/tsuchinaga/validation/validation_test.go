package validation

import (
	"reflect"
	"testing"
)

func TestNewValidator(t *testing.T) {
	t.Parallel()
	expect := &validator{}
	actual := NewValidator()
	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("%s 失敗\n期待: %+v\n実際: %+v\n", t.Name(), expect, actual)
	}
}

func TestValidator_IsValidDir(t *testing.T) {
	t.Parallel()
	testTable := []struct {
		desc   string
		arg    string
		expect bool
	}{
		{desc: "dirは必須なので空文字ならfalse", arg: "", expect: false},
		{desc: "ここではディレクトリかどうかまでは見ないので、空文字でなければtrue", arg: "foo", expect: true},
	}

	for _, test := range testTable {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			v := &validator{}
			actual := v.IsValidDir(test.arg)
			if test.expect != actual {
				t.Errorf("%s 失敗\n期待: %v\n実際: %v\n", t.Name(), test.expect, actual)
			}
		})
	}
}

func TestValidator_IsValidFileType(t *testing.T) {
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
			v := &validator{}
			actual := v.IsValidFileType(test.arg)
			if test.expect != actual {
				t.Errorf("%s 失敗\n期待: %v\n実際: %v\n", t.Name(), test.expect, actual)
			}
		})
	}
}

func TestValidator_IsValidSrc(t *testing.T) {
	t.Parallel()
	testTable := []struct {
		desc   string
		arg    string
		expect bool
	}{
		{desc: "許可されたファイルタイプならtrue", arg: "jpeg", expect: true},
		{desc: "許可されていないファイルタイプならfalse", arg: "gif", expect: false},
	}

	for _, test := range testTable {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			v := &validator{}
			actual := v.IsValidSrc(test.arg)
			if test.expect != actual {
				t.Errorf("%s 失敗\n期待: %v\n実際: %v\n", t.Name(), test.expect, actual)
			}
		})
	}
}

func TestValidator_IsValidDest(t *testing.T) {
	t.Parallel()
	testTable := []struct {
		desc   string
		arg1   string
		arg2   string
		expect bool
	}{
		{desc: "許可されたファイルタイプならtrue", arg1: "jpeg", arg2: "", expect: true},
		{desc: "許可されていないファイルタイプならfalse", arg1: "gif", arg2: "", expect: false},
		{desc: "許可されているファイルタイプでもsrcとdestが一致しているとfalse", arg1: "png", arg2: "png", expect: false},
	}

	for _, test := range testTable {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			v := &validator{}
			actual := v.IsValidDest(test.arg1, test.arg2)
			if test.expect != actual {
				t.Errorf("%s 失敗\n期待: %v\n実際: %v\n", t.Name(), test.expect, actual)
			}
		})
	}
}
