package imgconv

import (
	"image"
	"image/png"
	"io"
)

type PNG struct{}

func (*PNG) GetEncoder() Encoder {
	return &PNGEncoder{}
}

func (*PNG) Has(ext string) bool {
	var pngExt = map[string]bool{
		".png": true,
		".PNG": true,
	}

	return pngExt[ext]
}

type PNGEncoder struct{}

func (*PNGEncoder) execute(w io.Writer, Image image.Image) error {
	err := png.Encode(w, Image)
	return err
}
