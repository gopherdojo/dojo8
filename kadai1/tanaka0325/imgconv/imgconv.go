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

	// get target image paths from args
	udns := args.uniq()
	paths, err := getPaths(udns, *options.From)
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

func getPaths(dns []string, from string) ([]string, error) {
	paths := []string{}

	for _, dn := range dns {
		if err := filepath.Walk(dn, func(path string, info os.FileInfo, err error) error {
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
