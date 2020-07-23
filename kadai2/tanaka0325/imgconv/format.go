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

type ImageConverter interface {
	Decoder
	Encoder
	GetExt() string
}

type Image struct {
	Ext string
}

// ImagePng is type for png format.
type PNG Image

func (PNG) Decode(r io.Reader) (image.Image, error) { return png.Decode(r) }
func (PNG) Encode(w io.Writer, i image.Image) error { return png.Encode(w, i) }
func (p *PNG) GetExt() string                       { return p.Ext }

// JPEG is type for jpeg format.
type JPEG Image

func (JPEG) Decode(r io.Reader) (image.Image, error) { return jpeg.Decode(r) }
func (JPEG) Encode(w io.Writer, i image.Image) error { return jpeg.Encode(w, i, nil) }
func (j *JPEG) GetExt() string                       { return j.Ext }

// GIF is type for gif format.
type GIF Image

func (GIF) Decode(r io.Reader) (image.Image, error) { return gif.Decode(r) }
func (GIF) Encode(w io.Writer, i image.Image) error {
	return gif.Encode(w, i, &gif.Options{NumColors: 256})
}
func (g *GIF) GetExt() string { return g.Ext }

// BMP is type for bmp format.
type BMP Image

func (BMP) Decode(r io.Reader) (image.Image, error) { return bmp.Decode(r) }
func (BMP) Encode(w io.Writer, i image.Image) error { return bmp.Encode(w, i) }
func (b *BMP) GetExt() string                       { return b.Ext }

// TIFF is type for tiff format.
type TIFF Image

func (TIFF) Decode(r io.Reader) (image.Image, error) { return tiff.Decode(r) }
func (TIFF) Encode(w io.Writer, i image.Image) error { return tiff.Encode(w, i, nil) }
func (t *TIFF) GetExt() string                       { return t.Ext }

func NewImage(ext string) ImageConverter {
	switch ext {
	case "png":
		return &PNG{Ext: ext}
	case "jpg", "jpeg":
		return &JPEG{Ext: ext}
	case "gif":
		return &GIF{Ext: ext}
	case "bmp":
		return &BMP{Ext: ext}
	case "tiff", "tif":
		return &TIFF{Ext: ext}
	}

	return nil
}
