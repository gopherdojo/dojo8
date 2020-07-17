package imgconv

import (
	"image"
	"image/png"
	"io"
)

type PngImage struct{}

func (*PngImage) GetEncoder() Encoder {
	return &PngEncoder{}
}

func (*PngImage) IsMatchExt(ext string) bool {
	var pngExt = map[string]bool{
		".png": true,
		".PNG": true,
	}

	return pngExt[ext]
}

type PngEncoder struct{}

func (*PngEncoder) execute(w io.Writer, Image image.Image) error {
	err := png.Encode(w, Image)
	return err
}
