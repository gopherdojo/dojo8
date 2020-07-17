package imgconv

import (
	"fmt"
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
	fmt.Println("hoge")
	img, _, err := image.Decode(in)
	if err != nil {
		return err
	}

	fmt.Println("hoge")
	return c.encoder.execute(out, img)
}
