// Imgconv package is to convert images file format.
package imgconv

import "io"

// ConvertParam is parameter to convert image format.
type ConvertParam struct {
	Path         string
	FileHandler  FileHandler
	BeforeFormat ImageFormater
	AfterFormat  ImageFormater
}

// Do is func to convert image format.
func Do(param ConvertParam) (rerr error) {
	r, err := param.FileHandler.Open(param.Path)
	if err != nil {
		return err
	}
	defer r.Close()

	n := buildAfterPath(param.Path, param.BeforeFormat.GetExt(), param.AfterFormat.GetExt())
	w, err := param.FileHandler.Create(n)
	if err != nil {
		return err
	}
	defer func() {
		err := w.Close()
		if err != nil {
			rerr = err
		}
	}()

	if err := convert(r, param.BeforeFormat, w, param.AfterFormat); err != nil {
		return err
	}

	return nil
}

func buildAfterPath(path, beforeExt, afterExt string) string {
	e := len(path) - len(beforeExt)
	return path[:e] + afterExt
}

func convert(r io.Reader, d Decoder, w io.Writer, e Encoder) error {
	img, err := d.Decode(r)
	if err != nil {
		return err
	}

	if err := e.Encode(w, img); err != nil {
		return err
	}

	return nil
}
