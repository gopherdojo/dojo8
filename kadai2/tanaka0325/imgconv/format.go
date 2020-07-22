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

type ImageFormater interface {
	Decoder
	Encoder
	GetExt() string
}

type ImageFormat struct {
	Ext string
}

// ImagePng is type for png format.
type PNG ImageFormat

func (PNG) Decode(r io.Reader) (image.Image, error) { return png.Decode(r) }
func (PNG) Encode(w io.Writer, i image.Image) error { return png.Encode(w, i) }
func (p *PNG) GetExt() string                       { return p.Ext }

// JPEG is type for jpeg format.
type JPEG ImageFormat

func (JPEG) Decode(r io.Reader) (image.Image, error) { return jpeg.Decode(r) }
func (JPEG) Encode(w io.Writer, i image.Image) error { return jpeg.Encode(w, i, nil) }
func (j *JPEG) GetExt() string                       { return j.Ext }

// GIF is type for gif format.
type GIF ImageFormat

func (GIF) Decode(r io.Reader) (image.Image, error) { return gif.Decode(r) }
func (GIF) Encode(w io.Writer, i image.Image) error {
	return gif.Encode(w, i, &gif.Options{NumColors: 256})
}
func (g *GIF) GetExt() string { return g.Ext }

// BMP is type for bmp format.
type BMP ImageFormat

func (BMP) Decode(r io.Reader) (image.Image, error) { return bmp.Decode(r) }
func (BMP) Encode(w io.Writer, i image.Image) error { return bmp.Encode(w, i) }
func (b *BMP) GetExt() string                       { return b.Ext }

// TIFF is type for tiff format.
type TIFF ImageFormat

func (TIFF) Decode(r io.Reader) (image.Image, error) { return tiff.Decode(r) }
func (TIFF) Encode(w io.Writer, i image.Image) error { return tiff.Encode(w, i, nil) }
func (t *TIFF) GetExt() string                       { return t.Ext }

func NewImageFormat(ext string) ImageFormater {
	switch ext {
	case "png":
		return &PNG{Ext: "png"}
	case "jpg", "jpeg":
		return &JPEG{Ext: "jpeg"}
	case "gif":
		return &GIF{Ext: "gif"}
	case "bmp":
		return &BMP{Ext: "bmp"}
	case "tiff", "tif":
		return &TIFF{Ext: "tiff"}
	}

	return nil
}
