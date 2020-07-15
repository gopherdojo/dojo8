package imgconv

import "errors"

type exts []string

func (es exts) include(w string) error {
	for _, e := range es {
		if e == w {
			return nil
		}
	}
	return errors.New(w + " is not allowed")
}
