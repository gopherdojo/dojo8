// Imgconv package is to convert images file format.
package imgconv

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var allowedExts = []string{"png", "jpg", "jpeg", "gif", "bmp", "tiff", "tif"}

// Run is to convert image file format
func Run(options Options, args Args) error {
	// validator
	if err := options.validate(allowedExts); err != nil {
		return err
	}

	// get target image flepaths from args
	paths, err := getTargetFilePaths(args, *options.From)
	if err != nil {
		return err
	}

	// convert
	imgs, err := createConvImages(paths, *options.From, *options.To)
	if err != nil {
		return err
	}
	for _, img := range imgs {
		if err := img.decode(); err != nil {
			return err
		}

		if *options.DryRun {
			fmt.Println(img.filename+"."+img.fromExt, "=>", img.filename+"."+img.toExt)
		} else {
			if err := img.encode(); err != nil {
				return err
			}
		}
	}

	return nil
}

func getTargetFilePaths(args Args, from string) ([]string, error) {
	uns := args.uniq()

	paths := []string{}
	for _, n := range uns {
		b, err := isDir(n)
		if err != nil {
			return nil, err
		}
		if !b {
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

func createConvImages(paths []string, from, to string) ([]convImage, error) {
	images := []convImage{}
	for _, p := range paths {
		i := convImage{
			filename: strings.Replace(p, "."+from, "", 1),
			fromExt:  strings.ToLower(from),
			toExt:    strings.ToLower(to),
		}
		images = append(images, i)
	}

	return images, nil
}
