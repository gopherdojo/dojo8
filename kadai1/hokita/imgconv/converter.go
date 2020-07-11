package imgconv

import (
	"image"
	"os"
	"path/filepath"
)

func newConverter(from, to string) (*Converter, error) {
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

func (conv *Converter) Execute(path string) (rerr error) {
	// ignore unrelated file
	if !conv.fromImage.IsMatchExt(filepath.Ext(path)) {
		// TODO: os.IsNotExit対応する
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
	defer func() {
		if err := out.Close(); err != nil {
			rerr = err
		}
	}()
	if err != nil {
		return err
	}

	// output image
	if err := conv.toImage.Encode(out, img); err != nil {
		return err
	}
	return nil
}

func (conv *Converter) SwitchExt(path string) string {
	ext := filepath.Ext(path)
	toExt := conv.toImage.GetMainExt()

	return path[:len(path)-len(ext)] + toExt
}
