package imgconv

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

// convImage is type included infomation to convert file format.
type convImage struct {
	filename string
	fromExt  string
	toExt    string
	image    image.Image
}

func (i *convImage) decode() error {
	r, err := os.Open(i.filename + "." + i.fromExt)
	if err != nil {
		return err
	}
	defer r.Close()

	img, err := decodeHelper(r, i.fromExt)
	if err != nil {
		return fmt.Errorf("decode error: %w", err)
	}

	i.image = img
	return nil
}

func decodeHelper(r io.Reader, ext string) (image.Image, error) {
	switch ext {
	case "png":
		return png.Decode(r)
	case "jpg", "jpeg":
		return jpeg.Decode(r)
	case "gif":
		return gif.Decode(r)
	case "bmp":
		return bmp.Decode(r)
	case "tiff", "tif":
		return tiff.Decode(r)
	}
	return nil, fmt.Errorf("%s is not allowed", ext)
}

func (i *convImage) encode() error {
	w, err := os.Create(i.filename + "." + i.toExt)
	if err != nil {
		return err
	}
	defer func() error {
		if err := w.Close(); err != nil {
			return err
		}
		return nil
	}()

	switch i.toExt {
	case "png":
		return png.Encode(w, i.image)
	case "jpg", "jpeg":
		return jpeg.Encode(w, i.image, nil)
	case "gif":
		return gif.Encode(w, i.image, nil)
	case "bmp":
		return gif.Encode(w, i.image, nil)
	case "tiff", "tif":
		return gif.Encode(w, i.image, nil)
	}

	return fmt.Errorf("cannot encode image")
}
