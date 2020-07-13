package converter

import (
	"fmt"
	"testing"
)

func TestImageConverter(t *testing.T) {

	var tests = []struct {
		args    Args
		wantErr bool
		want    string
	}{
		{Args{FilePath: "../testdata/301_toJpg/_JR6xY.tiff", BeforeFormat: "tiff", AfterFormat: "jpg"},
			false, ""},
		{Args{FilePath: "../testdata/302_toPng/XvMdoh.bmp", BeforeFormat: "bmp", AfterFormat: "png"},
			false, ""},
		{Args{FilePath: "../testdata/303_toGif/gophercolor.png", BeforeFormat: "png", AfterFormat: "gif"},
			false, ""},
		{Args{FilePath: "../testdata/304_toBmp/XvMdoh.gif", BeforeFormat: "gif", AfterFormat: "bmp"},
			false, ""},
		{Args{FilePath: "../testdata/305_toTiff/vMfDls.jpg", BeforeFormat: "jpg", AfterFormat: "tiff"},
			false, ""},

		{Args{FilePath: "../testdata/391_notSupport/vMfDls.jpg", BeforeFormat: "jpg", AfterFormat: "HEIC"},
			true, "The specified image format is not supported. : ../testdata/391_notSupport/vMfDls.jpg"},
		{Args{FilePath: "../testdata/392_openError/vMfDls.jpg", BeforeFormat: "jpg", AfterFormat: "png"},
			true, "open ../testdata/392_openError/vMfDls.jpg: no such file or directory"},
		{Args{FilePath: "../testdata/393_decordError/notImage.jpg", BeforeFormat: "jpg", AfterFormat: "png"},
			true, "image: unknown format : ../testdata/393_decordError/notImage.jpg"},
		{Args{FilePath: "../testdata/394_createError/gophercolor.jpg", BeforeFormat: "jpg", AfterFormat: "png"},
			true, "open ../testdata/394_createError/gophercolor.png: permission denied : ../testdata/394_createError/gophercolor.jpg"},
		//{Args{FilePath: "../testdata/395_encordError", BeforeFormat: "jpg", AfterFormat: "png"},
		//	true, ""},
		//{Args{FilePath: "../testdata/396_closeError", BeforeFormat: "jpg", AfterFormat: "png"},
		//	true, ""},
	}

	for _, test := range tests {
		//失敗時にどのテストのどの引数で起きたのかを分かるようにする
		descr := fmt.Sprintf("ImageConverter(%v)", test.args)

		err := ImageConverter(test.args)

		switch {
		//エラー発生しないケースだが、エラーが発生した場合
		case !test.wantErr && err != nil:
			t.Errorf("want no err , %s = %q ", descr, err)
		//エラーが発生するケースだが、エラーが発生しない場合
		case test.wantErr && err == nil:
			t.Errorf("want err , %s = %q ", descr, err)
		//エラーが発生するケースで、エラー内容が異なる場合
		case test.wantErr && err.Error() != test.want:
			t.Errorf("%s = %q ,want %q", descr, err.Error(), test.want)

		}

	}
}
