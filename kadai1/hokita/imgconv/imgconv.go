package imgconv

import (
	"errors"
	"flag"
	"os"
	"path/filepath"
)

func Call() error {
	var (
		from = flag.String("from", "jpg", "Conversion source extension.")
		to   = flag.String("to", "png", "Conversion target extension.")
	)
	// TODO: mainパッケージへ移す
	flag.Parse()
	dir := flag.Arg(0)

	if flag.Arg(0) == "" {
		return errors.New("Please specify a directory.")
	}

	if f, err := os.Stat(dir); os.IsNotExist(err) || !f.IsDir() {
		return errors.New("Cannot find directory.")
	}

	converter, err := newConverter(*from, *to)
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
