package imgconv

import (
	"image"
	"image/png"
	"io"
)

type PngImage struct{}

var pngExt = map[string]bool{
	".png": true,
	".PNG": true,
}

func (PngImage) Encode(w io.Writer, m image.Image) error {
	err := png.Encode(w, m)
	return err
}

func (pi PngImage) IsMatchExt(ext string) bool {
	return pngExt[ext]
}

func (PngImage) GetMainExt() string {
	return ".png"
}
