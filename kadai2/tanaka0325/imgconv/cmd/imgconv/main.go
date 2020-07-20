package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gopherdojo/dojo8/kadai2/tanaka0325/imgconv"
)

var options imgconv.Options
var args []string
var allowedExts = []string{"png", "jpg", "jpeg", "gif", "bmp", "tiff", "tif"}

func init() {
	options.From = flag.String("f", "jpg", "file extension before convert")
	options.To = flag.String("t", "png", "file extension after convert")
	options.DryRun = flag.Bool("n", false, "dry run")
	flag.Parse()

	args = flag.Args()
}

func main() {
	// validate options
	if err := options.Validate(allowedExts); err != nil {
		onExit(err)
	}

	// get filenames
	dirnames := uniq(args)
	paths, err := getTargetFilenames(dirnames, *options.From)

	if err != nil {
		onExit(err)
	}

	// convert
	for _, path := range paths {
		param := imgconv.ConvertParam{
			File:        imgconv.NewFile(path),
			BeforeImage: imgconv.NewImage(*options.From),
			AfterImage:  imgconv.NewImage(*options.To),
			FromExt:     *options.From,
			ToExt:       *options.To,
		}

		if !*options.DryRun {
			if err := imgconv.Do(param); err != nil {
				onExit(err)
			}
		} else {
			fmt.Printf("%[1]s.%[2]s => %[1]s.%[3]s\n", param.File.Path, param.FromExt, param.ToExt)
		}
	}
}

func onExit(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func uniq([]string) []string {
	m := map[string]bool{}
	u := []string{}

	for _, v := range args {
		if !m[v] {
			m[v] = true

			u = append(u, v)
		}
	}

	return u
}

func getTargetFilenames(ds []string, from string) ([]string, error) {
	names := []string{}
	for _, n := range ds {
		if ok, err := isDir(n); err != nil {
			return nil, err
		} else if !ok {
			return nil, fmt.Errorf("%s is not a directory", n)
		}

		if err := filepath.Walk(n, func(name string, info os.FileInfo, err error) error {
			if filepath.Ext(name) == "."+from {
				n := strings.Replace(name, "."+from, "", -1)
				names = append(names, n)
			}
			return nil
		}); err != nil {
			return nil, err
		}
	}

	return names, nil
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
