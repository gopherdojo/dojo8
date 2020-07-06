package imgconv

import (
	"image"
	"image/jpeg"
	"io"
)

const QUALITY = 100

type JpegImage struct{}

func (_ JpegImage) Encode(w io.Writer, m image.Image) error {
	err := jpeg.Encode(w, m, &jpeg.Options{Quality: QUALITY})
	return err
}

func (ji JpegImage) IsMatchExt(ext string) bool {
	for _, myExt := range ji.Extensions() {
		if ext == myExt {
			return true
		}
	}
	return false
}

func (_ JpegImage) Extensions() []string {
	return []string{".jpg", ".jpeg", ".JPG", ".JPEG"}
}
