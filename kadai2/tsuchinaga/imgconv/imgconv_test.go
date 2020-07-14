package imgconv

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func init() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	if "imgconv" == filepath.Base(wd) {
		if err := os.Chdir(".."); err != nil {
			panic(err)
		}
	}
}

func TestConverter_IsDir(t *testing.T) {
	t.Parallel()
	testTable := []struct {
		desc    string
		arg     string
		expect1 bool
		expect2 error
	}{
		{desc: "ディレクトリを指定したらtrue", arg: "./testdata", expect1: true},
		{desc: "ファイルを指定したらfalse", arg: "./testdata/Example.jpg", expect1: false},
		{desc: "存在しないパスを指定したらエラー", arg: "./NotFound.jpg", expect2: FileStatError},
		{desc: "空文字を指定したらエラー", arg: "", expect2: FileStatError},
		// {desc: "不良セクタを指定したらfalse", arg1: "", expect1: false}, // やり方わからない
	}

	for _, test := range testTable {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			c := &converter{}
			actual1, actual2 := c.IsDir(test.arg)
			if test.expect1 != actual1 || !errors.Is(actual2, test.expect2) {
				t.Errorf("%s 失敗\n期待: %v, %v\n実際: %v, %v\n", t.Name(), test.expect1, test.expect2, actual1, actual2)
			}
		})
	}
}

func TestConverter_GetIMGType(t *testing.T) {
	t.Parallel()
	testTable := []struct {
		desc    string
		arg     string
		expect1 string
		expect2 error
	}{
		{desc: "jpegファイルを指定したらjpegが返される", arg: "./testdata/Example.jpg", expect1: "jpeg"},
		{desc: "pngファイルを指定したらpngが返される", arg: "./testdata/Example.png", expect1: "png"},
		{desc: "テキストファイルを指定したらエラー", arg: "./testdata/Example.txt", expect2: NotImageError},
		{desc: "画像ではないバイナリファイルを指定したらエラー", arg: "./testdata/Example.xlsx", expect2: NotImageError},
		{desc: "存在しないパスを指定したらエラー", arg: "./foo/bar", expect2: OpenFileError},
	}

	for _, test := range testTable {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			c := &converter{}
			actual1, actual2 := c.GetIMGType(test.arg)
			if test.expect1 != actual1 || !errors.Is(actual2, test.expect2) {
				t.Errorf("%s 失敗\n期待: %v, %v\n実際: %v, %v\n", t.Name(), test.expect1, test.expect2, actual1, actual2)
			}
		})
	}
}

func TestConverter_DirFileList(t *testing.T) {
	t.Parallel()
	testTable := []struct {
		desc    string
		arg     string
		expect1 []string
		expect2 []string
		expect3 error
	}{
		{desc: "存在しないパスを指定したらエラー", arg: "", expect3: ReadDirError},
		{desc: "存在するパスを指定したら直下にあるディレクトリとファイルが返される",
			arg:     "./testdata",
			expect1: []string{"testdata/subdir1", "testdata/subdir2"},
			expect2: []string{"testdata/Example.jpg", "testdata/Example.png", "testdata/Example.txt", "testdata/Example.xlsx"}},
	}

	for _, test := range testTable {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			c := &converter{}
			actual1, actual2, actual3 := c.DirFileList(test.arg)
			if !reflect.DeepEqual(test.expect1, actual1) || !reflect.DeepEqual(test.expect2, actual2) || !errors.Is(actual3, test.expect3) {
				t.Errorf("%s 失敗\n期待: %+v, %+v, %+v\n実際: %+v, %+v, %+v\n", t.Name(),
					test.expect1, test.expect2, test.expect3, actual1, actual2, actual3)
			}
		})
	}
}

func TestConverter_Convert(t *testing.T) {
	t.Parallel()
	testTable := []struct {
		desc   string
		arg1   string
		arg2   string
		arg3   string
		expect error
	}{
		{desc: "存在しないパスを指定したらエラー", arg1: "foo/bar", expect: OpenFileError},
		{desc: "画像じゃないファイルを指定したらエラー", arg1: "testdata/Example.txt", expect: NotImageError},
		{desc: "srcに指定した画像形式と違う画像形式なら何もせず終了", arg1: "testdata/Example.jpg", arg2: "png", expect: nil},
		// {desc: "Decodeできない画像ならエラー", arg1: "", expect: ReadImageError}, // 壊れた画像とか作ればいける？
		// {desc: "空っぽの変換後のファイルを生成出来なかったらエラー", arg1: "", expect: CreateDestinationFileError}, // 吐き出し先にロックされたファイルを作ればいける？
		// {desc: "jpgへの画像の変換に失敗したらエラー", arg1: "", expect: EncodeImageError}, // どうすれば失敗させられるのか分からない
		// {desc: "pngへの画像の変換に失敗したらエラー", arg1: "", expect: EncodeImageError}, // どうすれば失敗させられるのか分からない
		{desc: "正常に終了すればエラーなく終了", arg1: "testdata/subdir1/Example.jpg", arg2: "jpeg", arg3: "png", expect: nil},
	}

	for _, test := range testTable {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			c := &converter{}
			actual := c.Convert(test.arg1, test.arg2, test.arg3)
			if !errors.Is(actual, test.expect) {
				t.Errorf("%s 失敗\n期待: %+v\n実際: %+v\n", t.Name(), test.expect, actual)
			}
		})
	}
}

func TestNewConverter(t *testing.T) {
	expect := &converter{}
	actual := NewConverter()
	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("%s 失敗\n期待: %+v\n実際: %+v\n", t.Name(), expect, actual)
	}
}

func TestNewIMGConverter(t *testing.T) {
	expect := &imgConverter{converter: NewConverter()}
	actual := NewIMGConverter()
	if !reflect.DeepEqual(expect, actual) {
		t.Errorf("%s 失敗\n期待: %+v\n実際: %+v\n", t.Name(), expect, actual)
	}
}

type testConverter struct {
	dirFileListReturn map[int]struct {
		ret1 []string
		ret2 []string
		ret3 error
	}
	dirFileListCount int
	convertReturn    map[int]error
	convertCount     int
}

func (t testConverter) IsDir(string) (bool, error) {
	panic("implement me")
}

func (t testConverter) GetIMGType(string) (_ string, err error) {
	panic("implement me")
}

func (t *testConverter) DirFileList(string) ([]string, []string, error) {
	rs := t.dirFileListReturn[t.dirFileListCount]
	t.dirFileListCount++
	return rs.ret1, rs.ret2, rs.ret3
}

func (t *testConverter) Convert(string, string, string) (err error) {
	r := t.convertReturn[t.convertCount]
	t.convertCount++
	return r
}

func TestImgConverter_Do(t *testing.T) {
	t.Parallel()
	testTable := []struct {
		desc              string
		dirFileListReturn map[int]struct {
			ret1 []string
			ret2 []string
			ret3 error
		}
		convertReturn          map[int]error
		expectDirFileListCount int
		expectConvertCount     int
		expectErrorCount       int
	}{
		{desc: "dirが読み取れなければエラーを1回返して終わり",
			dirFileListReturn: map[int]struct {
				ret1 []string
				ret2 []string
				ret3 error
			}{0: {ret3: ReadDirError}},
			expectDirFileListCount: 1, expectConvertCount: 0, expectErrorCount: 1,
		},
		{desc: "dirが3段あれば3回実行する",
			dirFileListReturn: map[int]struct {
				ret1 []string
				ret2 []string
				ret3 error
			}{0: {ret1: []string{"subdir1"}, ret2: []string{}},
				1: {ret1: []string{"subdir2"}, ret2: []string{}},
				2: {ret1: []string{}, ret2: []string{}}},
			expectDirFileListCount: 3, expectConvertCount: 0, expectErrorCount: 0,
		},
		{desc: "fileの数だけconvertが叩かれる",
			dirFileListReturn: map[int]struct {
				ret1 []string
				ret2 []string
				ret3 error
			}{0: {ret1: []string{}, ret2: []string{"file1", "file2", "file3", "file4"}}},
			expectDirFileListCount: 1, expectConvertCount: 4, expectErrorCount: 0,
		},
		{desc: "fileの変換に失敗したらエラーが返される",
			dirFileListReturn: map[int]struct {
				ret1 []string
				ret2 []string
				ret3 error
			}{0: {ret1: []string{}, ret2: []string{"file1", "file2", "file3", "file4"}}},
			convertReturn:          map[int]error{0: ReadImageError, 2: ReadImageError},
			expectDirFileListCount: 1, expectConvertCount: 4, expectErrorCount: 2,
		},
		{desc: "エラーなくdirもfileもなければ何もせずに抜ける",
			dirFileListReturn: map[int]struct {
				ret1 []string
				ret2 []string
				ret3 error
			}{0: {ret1: []string{}, ret2: []string{}}},
			expectDirFileListCount: 1, expectConvertCount: 0, expectErrorCount: 0,
		},
	}

	for _, test := range testTable {
		test := test
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			testConverter := &testConverter{
				dirFileListReturn: test.dirFileListReturn,
				convertReturn:     test.convertReturn,
			}
			imgConverter := &imgConverter{converter: testConverter}

			var errCount int
			ch := imgConverter.Do("", "", "")
			for range ch {
				errCount++
			}

			if test.expectDirFileListCount != testConverter.dirFileListCount ||
				test.expectConvertCount != testConverter.convertCount ||
				test.expectErrorCount != errCount {
				t.Errorf("%s 失敗\n期待: %d, %d, %d\n実際: %d, %d, %d\n", t.Name(),
					test.expectDirFileListCount, test.expectConvertCount, test.expectErrorCount,
					testConverter.dirFileListCount, testConverter.convertCount, errCount)
			}
		})
	}
}
