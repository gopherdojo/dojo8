package conv

import "testing"

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
