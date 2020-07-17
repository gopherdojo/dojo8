package imgconv

import (
	"image"
	"image/jpeg"
	"io"
)

type JpegImage struct{}

func (*JpegImage) GetEncoder() Encoder {
	return &JpegEncoder{}
}

func (*JpegImage) IsMatchExt(ext string) bool {
	var jpegExt = map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".JPG":  true,
		".JPEG": true,
	}

	return jpegExt[ext]
}

type JpegEncoder struct{}

func (*JpegEncoder) execute(w io.Writer, Image image.Image) error {
	err := jpeg.Encode(w, Image, &jpeg.Options{Quality: jpeg.DefaultQuality})
	return err
}
