// Imgconv package is to convert images file format.
package imgconv

type ConvertParam struct {
	File        File
	BeforeImage Decoder
	AfterImage  Encoder
	FromExt     string
	ToExt       string
}

func Do(param ConvertParam) (err error) {
	// open file
	r, err := param.File.Open()
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
	w, err := param.File.Create()
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
