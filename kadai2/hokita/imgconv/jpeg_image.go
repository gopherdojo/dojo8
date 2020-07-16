package imgconv

import (
	"image"
	"image/jpeg"
	"io"
)

const QUALITY = 100

type JpegImage struct{}

var jpegExt = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".JPG":  true,
	".JPEG": true,
}

func (JpegImage) Encode(w io.Writer, m image.Image) error {
	err := jpeg.Encode(w, m, &jpeg.Options{Quality: QUALITY})
	return err
}

func (ji JpegImage) IsMatchExt(ext string) bool {
	return jpegExt[ext]
}

func (JpegImage) GetMainExt() string {
	return ".jpg"
}
