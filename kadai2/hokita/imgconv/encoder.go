package imgconv

import (
	"image"
	"io"
)

type Encoder interface {
	execute(w io.Writer, Image image.Image) error
}
