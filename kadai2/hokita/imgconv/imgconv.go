package imgconv

import (
	"image"
	"io"
)

type Converter struct {
	encoder Encoder
}

func NewConverter(encoder Encoder) *Converter {
	return &Converter{encoder}
}

func (c *Converter) Execute(in io.Reader, out io.Writer) error {
	img, _, err := image.Decode(in)
	if err != nil {
		return err
	}

	return c.encoder.execute(out, img)
}
