// Imgconv package is to convert images file format.
package imgconv

// ConvertParam is parameter to convert image format.
type ConvertParam struct {
	Path        string
	File        OpenCreator
	BeforeImage Decoder
	AfterImage  Encoder
	FromExt     string
	ToExt       string
}

// Do is func to convert image format.
func Do(param ConvertParam) (err error) {
	r, err := param.File.Open(param.Path)
	if err != nil {
		return
	}
	defer r.Close()

	img, err := param.BeforeImage.Decode(r)
	if err != nil {
		return
	}

	e := len(param.Path) - len(param.FromExt)
	w, err := param.File.Create(param.Path[:e] + param.ToExt)
	if err != nil {
		return err
	}

	defer func() {
		err = w.Close()
	}()

	if err := param.AfterImage.Encode(w, img); err != nil {
		return err
	}

	return
}
