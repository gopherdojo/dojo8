package main

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo8/kadai2/hokita/imgconv/imgconv"
)

func convert(dir, from, to string) error {
	if err := validation(dir, from, to); err != nil {
		return err
	}

	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) (rerr error) {
			fromImage, err := imgconv.NewImage("." + from)
			if err != nil {
				return err
			}

			toImage, err := imgconv.NewImage("." + to)
			if err != nil {
				return err
			}

			// ignore unrelated file
			if !fromImage.IsMatchExt(filepath.Ext(path)) {
				return nil
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			out, err := os.Create(switchExt(path, "."+to))
			defer func() {
				if err := out.Close(); err != nil {
					rerr = err
				}
			}()
			if err != nil {
				return err
			}

			converter := imgconv.NewConverter(toImage.GetEncoder())
			return converter.Execute(file, out)
		})
	if err != nil {
		return err
	}

	return nil

}

func switchExt(path, to string) string {
	ext := filepath.Ext(path)
	return path[:len(path)-len(ext)] + to
}

func validation(dir, from, to string) error {
	if dir == "" {
		return errors.New("please specify a directory")
	}

	if f, err := os.Stat(dir); os.IsNotExist(err) || !f.IsDir() {
		return errors.New("cannot find directory")
	}

	return nil
}
