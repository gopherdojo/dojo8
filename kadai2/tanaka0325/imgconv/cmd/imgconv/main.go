package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gopherdojo/dojo8/kadai2/tanaka0325/imgconv"
)

var args []string

var (
	from   string
	to     string
	dryRun bool
)

var allowedExts = map[string]bool{
	"png":  true,
	"jpg":  true,
	"jpeg": true,
	"gif":  true,
	"bmp":  true,
	"tiff": true,
	"tif":  true,
}

func init() {
	flag.StringVar(&from, "from", "jpg", "before ext")
	flag.StringVar(&from, "f", "jpg", "before ext (short)")
	flag.StringVar(&to, "to", "png", "after ext")
	flag.StringVar(&to, "t", "png", "after ext (short)")
	flag.BoolVar(&dryRun, "dry-run", false, "use dry-run")
	flag.BoolVar(&dryRun, "n", false, "use dry-run (short)")
	flag.Parse()

	args = flag.Args()
}

func main() {
	// validate options
	if ok := isAllowedFileType(from); !ok {
		onExit(fmt.Errorf("%s is invalid filetype", from))
	}
	if ok := isAllowedFileType(to); !ok {
		onExit(fmt.Errorf("%s is invalid filetype", to))
	}

	// validate arguments
	dirnames := uniq(args)
	for _, dirname := range dirnames {
		ok, err := isDir(dirname)
		if err != nil {
			onExit(err)
		}
		if !ok {
			onExit(fmt.Errorf("%s is not a directory", dirname))
		}
	}

	paths, err := getTargetFilepaths(dirnames, from)
	if err != nil {
		onExit(err)
	}

	for _, path := range paths {
		param := imgconv.ConvertParam{
			Path:        path,
			File:        imgconv.NewFile(),
			BeforeImage: imgconv.NewImage(from),
			AfterImage:  imgconv.NewImage(to),
			FromExt:     from,
			ToExt:       to,
		}

		if !dryRun {
			if err := imgconv.Do(param); err != nil {
				onExit(err)
			}
		} else {
			fmt.Printf("%[1]s.%[2]s => %[1]s.%[3]s\n", path, from, param.ToExt)
		}
	}
}

func isAllowedFileType(ft string) bool {
	return allowedExts[ft]
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

func getTargetFilepaths(ds []string, ext string) ([]string, error) {
	var names []string

	for _, n := range ds {
		if err := filepath.Walk(n, func(name string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if filepath.Ext(name) == "."+ext {
				n := strings.Replace(name, "."+ext, "", -1)
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
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return false, err
	}

	return fi.IsDir(), nil
}
