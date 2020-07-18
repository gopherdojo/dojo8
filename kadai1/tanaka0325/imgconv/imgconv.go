// Imgconv package is to convert images file format.
package imgconv

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var allowedExts = []string{"png", "jpg", "jpeg", "gif", "bmp", "tiff", "tif"}
var fromExt string
var toExt string

// Run is to convert image file format.
func Run(options Options, args Args) error {
	// // validator
	if err := options.validate(allowedExts); err != nil {
		return err
	}

	fromExt = *options.From
	toExt = *options.To

	// get target image flepaths from args
	paths, err := getTargetFilePaths(args, fromExt)
	if err != nil {
		return err
	}

	// convert
	for _, path := range paths {
		filename := strings.Replace(path, "."+fromExt, "", -1)
		// "))
		f := newCnvImage(fromExt)
		t := newCnvImage(toExt)

		if *options.DryRun {
			fmt.Printf("%s.%s => %s.%s \n", filename, fromExt, filename, toExt)
		} else if err := convert(f, t, filename); err != nil {
			return err
		}
	}

	return nil
}

func convert(d Decoder, e Encoder, filename string) (err error) {
	// open file
	r, err := os.Open(filename + "." + fromExt)
	if err != nil {
		return
	}
	defer r.Close()

	// decode
	img, err := d.Decode(r)
	if err != nil {
		return
	}

	// create file
	w, err := os.Create(filename + "." + toExt)
	if err != nil {
		return err
	}

	defer func() {
		err = w.Close()
	}()

	// encode
	if err := e.Encode(w, img); err != nil {
		return err
	}

	return
}

func getTargetFilePaths(args Args, from string) ([]string, error) {
	uns := args.uniq()

	paths := []string{}
	for _, n := range uns {
		if ok, err := isDir(n); err != nil {
			return nil, err
		} else if !ok {
			return nil, fmt.Errorf("%s is not a directory", n)
		}

		if err := filepath.Walk(n, func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == "."+from {
				paths = append(paths, path)
			}
			return nil
		}); err != nil {
			return nil, err
		}
	}

	return paths, nil
}

func isDir(path string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}

	fi, err := f.Stat()
	if err != nil {
		return false, err
	}

	return fi.IsDir(), nil
}
