package main

import (
	"bytes"
	"flag"
	"fmt"
	"testing"
)

func TestRun(t *testing.T) {
	var tests = []struct {
		bf      string
		af      string
		dir     string
		wantErr bool
		want    string
	}{
		{"tiff", "jpg", "./testdata/101_success", false, ""},
		{"jpg", "png", "./testdata/192_noDir", true, "stat ./testdata/192_noDir: no such file or directory\n"},
	}

	//エラー内容の捕捉をするため、出力先をテスト中変更
	outErr = new(bytes.Buffer)

	for _, test := range tests {
		//失敗時にどのテストのどの引数で起きたのかを分かるようにする
		descr := fmt.Sprintf("run() -bf=%s -af=%s -dir=%s", test.bf, test.af, test.dir)

		flag.CommandLine.Set("bf", test.bf)
		flag.CommandLine.Set("af", test.af)
		flag.CommandLine.Set("dir", test.dir)

		exitCode := run()
		var got string
		got = outErr.(*bytes.Buffer).String()

		switch {
		//エラー発生しないケースだが、エラーが発生した場合
		case !test.wantErr && exitCode != exitCodeSuccess:
			t.Errorf("want no err , %s = %q ", descr, got)
		//エラーが発生するケースだが、エラーが発生しない場合
		case test.wantErr && exitCode == exitCodeSuccess:
			t.Errorf("want err , %s = %q ", descr, got)
		//エラーが発生するケースで、エラー内容が異なる場合
		case test.wantErr && got != test.want:
			t.Errorf("%s = %q ,want %q", descr, got, test.want)

		}
	}
}

func TestCheckDir(t *testing.T) {
	var tests = []struct {
		dir     string
		wantErr bool
		want    string
	}{
		{"./testdata", false, ""},
		{"", true, "No directory specified"},
		{"./test", true, "stat ./test: no such file or directory"},
		{"./README.md", true, "stat ./README.md: no such file or directory"},
	}

	for _, test := range tests {
		//失敗時にどのテストのどの引数で起きたのかを分かるようにする
		descr := fmt.Sprintf("checkDir(%s)", test.dir)

		err := checkDir(test.dir)

		var got string
		if err != nil {
			got = err.Error()
		}

		errorfHelper(t, descr, test.wantErr, err, test.want, got)

	}
}

func TestConvert(t *testing.T) {
	var tests = []struct {
		bf      string
		af      string
		dir     string
		wantErr bool
		want    string
	}{
		{"TIFF", "JPG", "./testdata/201_lower", false, ""},
		{"tiff", "jpg", "./testdata/202_subDir", false, ""},
	}

	for _, test := range tests {
		//失敗時にどのテストのどの引数で起きたのかを分かるようにする
		descr := fmt.Sprintf("convert(%s, %s, %s)", test.af, test.bf, test.dir)

		err := convert(test.af, test.bf, test.dir)

		var got string
		if err != nil {
			got = err.Error()
		}

		errorfHelper(t, descr, test.wantErr, err, test.want, got)

	}
}

func errorfHelper(t *testing.T, descr string, wantErr bool, err error, want string, got string) {
	t.Helper()

	switch {
	//エラー発生しないケースだが、エラーが発生した場合
	case !wantErr && err != nil:
		t.Errorf("want no err , %s = %q ", descr, err)
	//エラーが発生するケースだが、エラーが発生しない場合
	case wantErr && err == nil:
		t.Errorf("want err , %s = %q ", descr, err)
	//エラーが発生するケースで、エラー内容が異なる場合
	case wantErr && got != want:
		t.Errorf("%s = %q ,want %q", descr, got, want)
	}
}
