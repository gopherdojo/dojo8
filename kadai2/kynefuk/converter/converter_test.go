package converter

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"testing"
)

var fileExtList = []string{
	"jpg",
	"jpeg",
	"png",
	"gif",
	"bmp",
	"tiff",
}

const testDirPath = "../testdata"

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
	default:
		log.Fatalf("failed to create test file. fileExt: %s", fileExt)
	}

	if err != nil {
		log.Fatalf("failed to create test file. fileExt: %s", fileExt)
	}

	return f
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func TestConverter(t *testing.T) {
	tmpDir := CreateTmpDir()
	defer os.RemoveAll(tmpDir)

	for _, from := range fileExtList {
		for _, to := range fileExtList {
			tmpFile := CreateTmpFile(tmpDir, from)
			fmt.Println("tmpFile:", tmpFile.Name())
			defer os.Remove(tmpFile.Name())
			imgConverter := NewConverter(from, to)
			err := imgConverter.ConvertFormat(tmpFile.Name(), ConvertExt(tmpFile.Name(), from, to))
			if err != nil {
				t.Errorf("failed to test. error: %s", err)
			}
			exists := Exists(ConvertExt(tmpFile.Name(), from, to))
			if !exists {
				t.Errorf("want true, but got %v", exists)
			}
		}
	}
}
