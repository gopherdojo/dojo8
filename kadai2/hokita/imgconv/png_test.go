package imgconv

import (
	"image"
	"os"
	"reflect"
	"testing"
)

func TestPNG_GetEncoder(t *testing.T) {
	test := struct {
		want Encoder
	}{
		want: &PNGEncoder{},
	}

	var pngImage PNG
	got := pngImage.GetEncoder()
	if !reflect.DeepEqual(got, test.want) {
		t.Errorf(`want="%v" got="%v"`, test.want, got)
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
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := out.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	img, _, err := image.Decode(file)
	if err != nil {
		t.Fatal(err)
	}

	var encoder PNGEncoder
	err = encoder.execute(out, img)
	if err != nil {
		t.Errorf("failed to call execute(): %v", err)
	}

	checkAndDeleteFile(t, outFile)
}
