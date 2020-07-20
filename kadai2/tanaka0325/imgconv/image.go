package imgconv

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

type Decoder interface {
	Decode(io.Reader) (image.Image, error)
}

type Encoder interface {
	Encode(io.Writer, image.Image) error
}

type DecodeEncoder interface {
	Decoder
	Encoder
}

// ImagePng is type for png format.
type PNG struct{}

func (PNG) Decode(r io.Reader) (image.Image, error) { return png.Decode(r) }

func (PNG) Encode(w io.Writer, i image.Image) error { return png.Encode(w, i) }

// JPEG is type for jpeg format.
type JPEG struct{}

func (JPEG) Decode(r io.Reader) (image.Image, error) { return jpeg.Decode(r) }

func (JPEG) Encode(w io.Writer, i image.Image) error { return jpeg.Encode(w, i, nil) }

// GIF is type for gif format.
type GIF struct{}

func (GIF) Decode(r io.Reader) (image.Image, error) { return gif.Decode(r) }

func (GIF) Encode(w io.Writer, i image.Image) error {
	return gif.Encode(w, i, &gif.Options{NumColors: 256})
}

// BMP is type for bmp format.
type BMP struct{}

func (BMP) Decode(r io.Reader) (image.Image, error) { return bmp.Decode(r) }

func (BMP) Encode(w io.Writer, i image.Image) error { return bmp.Encode(w, i) }

// TIFF is type for tiff format.
type TIFF struct{}

func (TIFF) Decode(r io.Reader) (image.Image, error) { return tiff.Decode(r) }

func (TIFF) Encode(w io.Writer, i image.Image) error { return tiff.Encode(w, i, nil) }

func NewImage(ext string) DecodeEncoder {
	switch ext {
	case "png":
		return PNG{}
	case "jpg", "jpeg":
		return JPEG{}
	case "gif":
		return GIF{}
	case "bmp":
		return BMP{}
	case "tiff", "tif":
		return TIFF{}
	}

	return nil
}
