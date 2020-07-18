package imgconv

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// ImageExtension represents valid image extensions.
type ImageExtension string

const (
	// GIF is the extension like foo.gif.
	GIF ImageExtension = "gif"
	// JPEG is the extension like foo.jpeg.
	JPEG ImageExtension = "jpeg"
	// JPG is the extension like foo.jpg.
	JPG ImageExtension = "jpg"
	// PNG is the extension like foo.png.
	PNG ImageExtension = "png"
)

var (
	// ValidImageExtensions is a list of the valid image extensions.
	ValidImageExtensions = []ImageExtension{GIF, JPEG, JPG, PNG}
)

func (i ImageExtension) String() string {
	return string(i)
}

// StringToImageExtension converts the string to an ImageExtension.
func StringToImageExtension(str string) (*ImageExtension, error) {
	if str[0] == '.' {
		str = str[1:]
	}

	for _, v := range ValidImageExtensions {
		if v.String() == strings.ToLower(str) {
			return &v, nil
		}
	}

	return nil, fmt.Errorf("invalid image extension: %s", str)
}

// ImageConverter represents the image coverter.
type ImageConverter struct {
	// From is the image extension before conversion.
	From ImageExtension
	// To is the image extension after conversion.
	To ImageExtension
}

// NewImageConverter returns a new ImageConverter.
func NewImageConverter(from, to string) (*ImageConverter, error) {
	extFrom, err := StringToImageExtension(from)
	if err != nil {
		return nil, err
	}

	extTo, err := StringToImageExtension(to)
	if err != nil {
		return nil, err
	}

	c := &ImageConverter{From: *extFrom, To: *extTo}

	return c, nil
}

func (i ImageConverter) readImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		err = f.Close()
	}()

	m, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (i ImageConverter) saveImage(m image.Image, path string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
	}()

	switch i.To {
	case GIF:
		err = gif.Encode(f, m, nil)
	case JPEG, JPG:
		err = jpeg.Encode(f, m, nil)
	case PNG:
		err = png.Encode(f, m)
	default:
		err = fmt.Errorf("invalid image format")
	}

	if err != nil {
		return err
	}

	return nil
}

// Convert convets the image extension from i.From to i.To.
func (i ImageConverter) Convert(src string, verbose bool) error {
	dir := filepath.Dir(src)
	dst := filepath.Join(dir, fmt.Sprintf("%s.%s", getFilename(src), i.To.String()))

	if _, err := os.Stat(dst); !os.IsNotExist(err) {
		return fmt.Errorf("file already exists: %s", dst)
	}

	img, err := i.readImage(src)
	if err != nil {
		return err
	}

	if err := i.saveImage(img, dst); err != nil {
		return err
	}

	if verbose {
		fmt.Printf("%s -> %s\n", src, dst)
	}

	return nil
}

// ConvertAll converts images in the directory.
func (i ImageConverter) ConvertAll(dir string, verbose bool) error {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path)[1:] == i.From.String() {
			if err := i.Convert(path, verbose); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func getFilename(path string) string {
	return filepath.Base(path[:len(path)-len(filepath.Ext(path))])
}
