package imgconv

import (
	"image"
	"os"
	"reflect"
	"testing"
)

func TestPngImage_GetEncoder(t *testing.T) {
	test := struct {
		want Encoder
	}{
		want: &PngEncoder{},
	}

	pngImage := PngImage{}
	got := pngImage.GetEncoder()
	if !reflect.DeepEqual(got, test.want) {
		t.Errorf(
			`want="%v" got="%v"`,
			test.want, got,
		)
	}
}

func TestPngEncoder_execute(t *testing.T) {
	inFile := "../testdata/test4/gopher.jpg"
	outFile := "../testdata/test4/gopher.png"

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

	encoder := &PngEncoder{}
	err = encoder.execute(out, img)
	if err != nil {
		t.Errorf("failed to call execute(): %v", err)
	}

	checkAndDeleteFile(t, outFile)
}
