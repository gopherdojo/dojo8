package imgconv

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var (
	// flags
	f      = flag.String("f", "jpg", "file extention before convert")
	t      = flag.String("t", "png", "file extention after convert")
	dryRun = flag.Bool("n", false, "dry run")

	// allow extensions
	allowedExts = exts{"png", "jpg", "jpeg", "gif", "bmp", "tiff", "tif"}
)

func init() {
	flag.Parse()
}

// Run is to convert image file format
func Run() {
	// check options ext
	to := strings.ToLower(*t)
	from := strings.ToLower(*f)
	targetExts := []string{to, from}
	for _, e := range targetExts {
		if err := allowedExts.include(e); err != nil {
			log.Fatal(fmt.Errorf("%w. ext is only allowd in %s", err, allowedExts))
		}
	}

	// get target image paths from args
	dns := flag.Args()
	udns := uniq(dns)
	paths, err := getPaths(udns, from)
	if err != nil {
		log.Fatal(err)
	}

	// convert
	imgs, err := createConvImages(paths, from, to)
	if err != nil {
		log.Fatal(err)
	}
	for _, img := range imgs {
		if err := img.decode(); err != nil {
			log.Fatal(err)
		}

		if *dryRun {
			fmt.Println(img.filename+"."+img.fromExt, "=>", img.filename+"."+img.toExt)
		} else {
			if err := img.encode(); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func uniq(s []string) []string {
	m := map[string]bool{}
	u := []string{}

	for _, v := range s {
		if !m[v] {
			m[v] = true
			u = append(u, v)
		}
	}

	return u
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
