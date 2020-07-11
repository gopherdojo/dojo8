package imgconv

import (
	"errors"
	"os"
	"path/filepath"
)

func Call(dir, from, to string) error {
	if dir == "" {
		return errors.New("Please specify a directory.")
	}

	if f, err := os.Stat(dir); os.IsNotExist(err) || !f.IsDir() {
		return errors.New("Cannot find directory.")
	}

	converter, err := newConverter(from, to)
	if err != nil {
		return err
	}

	err = filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			err = converter.Execute(path)
			return err
		})
	if err != nil {
		return err
	}

	return nil
}
