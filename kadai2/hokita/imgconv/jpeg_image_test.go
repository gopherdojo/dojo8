package imgconv

import (
	"image"
	"os"
	"reflect"
	"testing"
)

func TestJpegImage_GetEncoder(t *testing.T) {
	test := struct {
		want Encoder
	}{
		want: &JpegEncoder{},
	}

	jpegImage := JpegImage{}
	got := jpegImage.GetEncoder()
	if !reflect.DeepEqual(got, test.want) {
		t.Errorf(
			`want="%v" got="%v"`,
			test.want, got,
		)
	}
}

func TestJpegEncoder_execute(t *testing.T) {
	inFile := "../testdata/test3/gopher.png"
	outFile := "../testdata/test3/gopher.jpg"

	file, err := os.Open(inFile)
	if err != nil {
		t.Fatal(err)
	}

	out, err := os.Create(outFile)
	defer func() {
		if err := out.Close(); err != nil {
			t.Fatal(err)
		}
	}()
	if err != nil {
		t.Fatal(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		t.Fatal(err)
	}

	encoder := &JpegEncoder{}
	err = encoder.execute(out, img)
	if err != nil {
		t.Errorf("failed to call execute(): %v", err)
	}

	checkAndDeleteFile(t, outFile)
}
