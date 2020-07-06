package imgconv

import (
	"image"
	"os"
	"path/filepath"
)

func converterFactory(from string, to string) (*Converter, error) {
	fromImage, err := selectImage("." + from)
	if err != nil {
		return nil, err
	}

	toImage, err := selectImage("." + to)
	if err != nil {
		return nil, err
	}

	return &Converter{fromImage, toImage}, nil
}

type Converter struct {
	fromImage ImageType
	toImage   ImageType
}

func (conv *Converter) Execute(path string) error {
	// ignore unrelated file
	if !conv.fromImage.IsMatchExt(filepath.Ext(path)) {
		return nil
	}

	// file open
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		return err
	}

	// convert to image obj
	img, _, err := image.Decode(file)
	if err != nil {
		return err
	}

	// output file
	out, err := os.Create(conv.SwitchExt(path))
	defer out.Close()
	if err != nil {
		return err
	}

	// output image
	conv.toImage.Encode(out, img)
	return nil
}

func (conv *Converter) SwitchExt(path string) string {
	ext := filepath.Ext(path)
	toExt := conv.toImage.Extensions()[0]

	return path[:len(path)-len(ext)] + toExt
}
