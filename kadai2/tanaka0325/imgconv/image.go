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
type ImagePNG struct{}

func (ImagePNG) Decode(r io.Reader) (image.Image, error) { return png.Decode(r) }

func (ImagePNG) Encode(w io.Writer, i image.Image) error { return png.Encode(w, i) }

// ImageJPEG is type for jpeg format.
type ImageJPEG struct{}

func (ImageJPEG) Decode(r io.Reader) (image.Image, error) { return jpeg.Decode(r) }

func (ImageJPEG) Encode(w io.Writer, i image.Image) error { return jpeg.Encode(w, i, nil) }

// ImageGIF is type for gif format.
type ImageGIF struct{}

func (ImageGIF) Decode(r io.Reader) (image.Image, error) { return gif.Decode(r) }

func (ImageGIF) Encode(w io.Writer, i image.Image) error {
	return gif.Encode(w, i, &gif.Options{NumColors: 256})
}

// ImageBMP is type for bmp format.
type ImageBMP struct{}

func (ImageBMP) Decode(r io.Reader) (image.Image, error) { return bmp.Decode(r) }

func (ImageBMP) Encode(w io.Writer, i image.Image) error { return bmp.Encode(w, i) }

// ImageTIFF is type for tiff format.
type ImageTIFF struct{}

func (ImageTIFF) Decode(r io.Reader) (image.Image, error) { return tiff.Decode(r) }

func (ImageTIFF) Encode(w io.Writer, i image.Image) error { return tiff.Encode(w, i, nil) }

func NewImage(ext string) DecodeEncoder {
	switch ext {
	case "png":
		return ImagePNG{}
	case "jpg", "jpeg":
		return ImageJPEG{}
	case "gif":
		return ImageGIF{}
	case "bmp":
		return ImageBMP{}
	case "tiff", "tif":
		return ImageTIFF{}
	}

	return nil
}
