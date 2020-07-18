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

type CnvImage struct{}

// CnvImagePng is type for png format.
type CnvImagePNG CnvImage

func (ip CnvImagePNG) Decode(r io.Reader) (image.Image, error) { return png.Decode(r) }

func (ip CnvImagePNG) Encode(w io.Writer, i image.Image) error { return png.Encode(w, i) }

// CnvImageJPEG is type for jpeg format.
type CnvImageJPEG CnvImage

func (ip CnvImageJPEG) Decode(r io.Reader) (image.Image, error) { return jpeg.Decode(r) }

func (ip CnvImageJPEG) Encode(w io.Writer, i image.Image) error { return jpeg.Encode(w, i, nil) }

// CnvImageGIF is type for gif format.
type CnvImageGIF CnvImage

func (ip CnvImageGIF) Decode(r io.Reader) (image.Image, error) { return gif.Decode(r) }

func (ip CnvImageGIF) Encode(w io.Writer, i image.Image) error {
	return gif.Encode(w, i, &gif.Options{NumColors: 256})
}

// CnvImageBMP is type for bmp format.
type CnvImageBMP CnvImage

func (ip CnvImageBMP) Decode(r io.Reader) (image.Image, error) { return bmp.Decode(r) }

func (ip CnvImageBMP) Encode(w io.Writer, i image.Image) error { return bmp.Encode(w, i) }

// CnvImageTIFF is type for tiff format.
type CnvImageTIFF CnvImage

func (ip CnvImageTIFF) Decode(r io.Reader) (image.Image, error) { return tiff.Decode(r) }

func (ip CnvImageTIFF) Encode(w io.Writer, i image.Image) error { return tiff.Encode(w, i, nil) }

func NewCnvImage(ext string) DecodeEncoder {
	switch ext {
	case "png":
		return &CnvImagePNG{}
	case "jpg", "jpeg":
		return &CnvImageJPEG{}
	case "gif":
		return &CnvImageGIF{}
	case "bmp":
		return &CnvImageBMP{}
	case "tiff", "tif":
		return &CnvImageTIFF{}
	}

	return nil
}
