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

type Image struct{}

// ImagePng is type for png format.
type ImagePNG Image

func (ip ImagePNG) Decode(r io.Reader) (image.Image, error) { return png.Decode(r) }

func (ip ImagePNG) Encode(w io.Writer, i image.Image) error { return png.Encode(w, i) }

// ImageJPEG is type for jpeg format.
type ImageJPEG Image

func (ip ImageJPEG) Decode(r io.Reader) (image.Image, error) { return jpeg.Decode(r) }

func (ip ImageJPEG) Encode(w io.Writer, i image.Image) error { return jpeg.Encode(w, i, nil) }

// ImageGIF is type for gif format.
type ImageGIF Image

func (ip ImageGIF) Decode(r io.Reader) (image.Image, error) { return gif.Decode(r) }

func (ip ImageGIF) Encode(w io.Writer, i image.Image) error {
	return gif.Encode(w, i, &gif.Options{NumColors: 256})
}

// ImageBMP is type for bmp format.
type ImageBMP Image

func (ip ImageBMP) Decode(r io.Reader) (image.Image, error) { return bmp.Decode(r) }

func (ip ImageBMP) Encode(w io.Writer, i image.Image) error { return bmp.Encode(w, i) }

// ImageTIFF is type for tiff format.
type ImageTIFF Image

func (ip ImageTIFF) Decode(r io.Reader) (image.Image, error) { return tiff.Decode(r) }

func (ip ImageTIFF) Encode(w io.Writer, i image.Image) error { return tiff.Encode(w, i, nil) }

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
