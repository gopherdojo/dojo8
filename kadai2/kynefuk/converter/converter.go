package converter

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

// Converter converts image file
type Converter struct {
	From, To string
}

// ConvertFormat converts file format
func (converter *Converter) ConvertFormat(filepath, dst string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("failed to open file. file: %s", filepath)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("failed to decode file. file: %s", filepath)
	}

	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create output file. file: %s", dst)
	}
	defer out.Close()

	switch converter.To {
	case "jpg", "jpeg":
		err = jpeg.Encode(out, img, &jpeg.Options{Quality: 100})
	case "png":
		err = png.Encode(out, img)
	case "gif":
		err = gif.Encode(out, img, &gif.Options{NumColors: 256, Quantizer: nil, Drawer: nil})
	case "bmp":
		err = bmp.Encode(out, img)
	case "tiff":
		err = tiff.Encode(out, img, nil)
	default:
		return fmt.Errorf("unknown format type")
	}

	return err
}

// ConvertExt converts file format extention
func ConvertExt(filepath, from, to string) string {
	return strings.Replace(filepath, from, to, 1)
}

// NewConverter is a Constructor of Converter
func NewConverter(from, to string) *Converter {
	return &Converter{From: from, To: to}
}
