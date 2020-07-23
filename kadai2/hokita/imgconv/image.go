package imgconv

import (
	"errors"
)

type Image interface {
	GetEncoder() Encoder
	Has(ext string) bool
}

func NewImage(ext string) (Image, error) {
	var jpeg JPEG
	var png PNG

	switch {
	case jpeg.Has(ext):
		return &jpeg, nil
	case png.Has(ext):
		return &png, nil
	}

	return nil, errors.New("selected extension is not supported")
}
