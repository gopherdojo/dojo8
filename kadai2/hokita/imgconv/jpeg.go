package imgconv

import (
	"image"
	"image/jpeg"
	"io"
)

type JPEG struct{}

func (*JPEG) GetEncoder() Encoder {
	return &JPEGEncoder{}
}

func (*JPEG) Has(ext string) bool {
	var jpegExt = map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".JPG":  true,
		".JPEG": true,
	}

	return jpegExt[ext]
}

type JPEGEncoder struct{}

func (*JPEGEncoder) execute(w io.Writer, Image image.Image) error {
	err := jpeg.Encode(w, Image, &jpeg.Options{Quality: jpeg.DefaultQuality})
	return err
}
