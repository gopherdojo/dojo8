package helper

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

var FileExtList = []string{
	"jpg",
	"jpeg",
	"png",
	"gif",
	"bmp",
	"tiff",
}

func CreateTmpDir() string {
	dir, err := ioutil.TempDir("./", "example")
	if err != nil {
		log.Fatalf("failed to create test dir. error: %s", err)
	}
	return dir
}

func CreateJPG(f *os.File, img *image.RGBA) error {
	return jpeg.Encode(f, img, &jpeg.Options{Quality: 100})
}

func CreatePNG(f *os.File, img *image.RGBA) error {
	return png.Encode(f, img)
}

func CreateGIF(f *os.File, img *image.RGBA) error {
	return gif.Encode(f, img, &gif.Options{NumColors: 256, Quantizer: nil, Drawer: nil})
}

func CreateBMP(f *os.File, img *image.RGBA) error {
	return bmp.Encode(f, img)
}

func CreateTIFF(f *os.File, img *image.RGBA) error {
	return tiff.Encode(f, img, nil)
}

func CreateTmpFile(dir, fileExt string) *os.File {

	img := image.NewRGBA(image.Rect(0, 0, 100, 50))

	f, err := ioutil.TempFile(dir, "example.*."+fileExt)
	if err != nil {
		log.Fatalf("failed to create tmp file. error: %s", err)
	}
	defer f.Close()

	switch fileExt {
	case "jpg", "jpeg":
		err = CreateJPG(f, img)
	case "png":
		err = CreatePNG(f, img)
	case "gif":
		err = CreateGIF(f, img)
	case "bmp":
		err = CreateBMP(f, img)
	case "tiff":
		err = CreateTIFF(f, img)
	}

	if err != nil {
		log.Fatalf("failed to create test file. fileExt: %s", fileExt)
	}

	return f
}

func CreateTmpFiles(dir string) []string {
	var tmpFilePaths []string
	for _, v := range FileExtList {
		f, err := ioutil.TempFile(dir, "example.*."+v)
		if err != nil {
			log.Fatalf("failed to create tmp file. error: %s", err)
		}
		tmpFilePaths = append(tmpFilePaths, f.Name())
	}
	return tmpFilePaths
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func ErrorHelper(tb testing.TB, want, got interface{}) {
	tb.Helper()
	tb.Errorf("want = %v, got = %v", want, got)
}
