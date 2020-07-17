package imgconv

import (
	"errors"
)

type Image interface {
	GetEncoder() Encoder
	IsMatchExt(ext string) bool
}

func NewImage(ext string) (Image, error) {
	jpegImage := &JpegImage{}
	pngImage := &PngImage{}

	switch {
	case jpegImage.IsMatchExt(ext):
		return jpegImage, nil
	case pngImage.IsMatchExt(ext):
		return pngImage, nil
	}

	return nil, errors.New("selected extension is not supported")
}
