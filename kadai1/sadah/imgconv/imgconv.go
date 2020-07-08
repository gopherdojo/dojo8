package imgconv

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

const (
	canNotOpenImageFile  = "Can not open an image file: %s"
	canNotWriteImageFile = "Can not write an image file: %s"
	unsupportedExtension = "Unsupported extension: %s"
)

type exts []string

var (
	// SupportedExts is a list of supported extensions
	SupportedExts exts = []string{
		".jpg", ".jpeg", ".JPG", ".JPEG",
		".png", ".PNG",
		".gif", ".GIF",
		".bmp", ".BMP",
		".tiff", ".TIFF",
	}
)

// VerifySupportedExt verifies that the extension is supported.
func VerifySupportedExt(ext string) error {
	var isValidExt bool

	for _, e := range SupportedExts {
		if e == ext {
			isValidExt = true
		}
	}
	if !isValidExt {
		return fmt.Errorf(unsupportedExtension, ext)
	}
	return nil
}

// Conv converts the image file to tarExt format.
func Conv(path, tarExt string) error {
	src := path
	dst := path[:len(path)-len(filepath.Ext(path))] + tarExt

	log.Printf("Source file: %s", src)
	log.Printf("Target file: %s", dst)

	sf, err := os.Open(src)
	if err != nil {
		return fmt.Errorf(canNotOpenImageFile, src)
	}
	defer sf.Close()

	df, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf(canNotWriteImageFile, dst)
	}
	defer df.Close()

	img, _, err := image.Decode(sf)
	if err != nil {
		return err
	}

	switch strings.ToLower(filepath.Ext(dst)) {
	case ".png":
		err = png.Encode(df, img)
	case ".jpeg", ".jpg":
		err = jpeg.Encode(df, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	case ".gif":
		err = gif.Encode(df, img, &gif.Options{NumColors: 256, Quantizer: nil, Drawer: nil})
	case ".bmp":
		err = bmp.Encode(df, img)
	case ".tiff":
		err = tiff.Encode(df, img, nil)
	}

	if err != nil {
		return fmt.Errorf(canNotWriteImageFile, dst)
	}

	return nil
}
