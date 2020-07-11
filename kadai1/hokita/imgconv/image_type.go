package imgconv

import (
	"errors"
	"image"
	"io"
)

type ImageType interface {
	Encode(w io.Writer, m image.Image) error
	IsMatchExt(ext string) bool
	GetMainExt() string
}

func selectImage(ext string) (ImageType, error) {
	pngImage := PngImage{}
	jpegImage := JpegImage{}

	if pngImage.IsMatchExt(ext) {
		return pngImage, nil
	} else if jpegImage.IsMatchExt(ext) {
		return jpegImage, nil
	}

	return nil, errors.New("Selected extension is not supported.")
}
