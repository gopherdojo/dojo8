// Imgconv package is to convert images file format.
package imgconv

import (
	"os"
)

type ConvertParam struct {
	Path        string
	BeforeImage Decoder
	AfterImage  Encoder
	FromExt     string
	ToExt       string
}

func Convert(param ConvertParam) (err error) {
	// open file
	r, err := os.Open(param.Path + "." + param.FromExt)
	if err != nil {
		return
	}
	defer r.Close()

	// decode
	img, err := param.BeforeImage.Decode(r)
	if err != nil {
		return
	}

	// create file
	w, err := os.Create(param.Path + "." + param.ToExt)
	if err != nil {
		return err
	}

	defer func() {
		err = w.Close()
	}()

	// encode
	if err := param.AfterImage.Encode(w, img); err != nil {
		return err
	}

	return
}
