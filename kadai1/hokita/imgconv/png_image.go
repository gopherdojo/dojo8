package imgconv

import (
	"image"
	"image/png"
	"io"
)

type PngImage struct{}

func (_ PngImage) Encode(w io.Writer, m image.Image) error {
	err := png.Encode(w, m)
	return err
}

func (pi PngImage) IsMatchExt(ext string) bool {
	for _, myExt := range pi.Extensions() {
		if ext == myExt {
			return true
		}
	}
	return false
}

func (_ PngImage) Extensions() []string {
	return []string{".png", ".PNG"}
}
