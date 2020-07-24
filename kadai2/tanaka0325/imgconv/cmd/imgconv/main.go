package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo8/kadai2/tanaka0325/imgconv"
)

var (
	args []string

	from   string
	to     string
	dryRun bool
)

const (
	fromUsageText   = "before extension"
	toUsageText     = "after extension"
	dryRunUsageText = "with dry-run"
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
	flag.StringVar(&from, "from", "jpg", fromUsageText)
	flag.StringVar(&from, "f", "jpg", fromUsageText+" (short)")
	flag.StringVar(&to, "to", "png", toUsageText)
	flag.StringVar(&to, "t", "png", toUsageText+" (short)")
	flag.BoolVar(&dryRun, "dry-run", false, dryRunUsageText)
	flag.BoolVar(&dryRun, "n", false, dryRunUsageText+" (short)")
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
			Path:         path,
			FileHandler:  imgconv.NewFile(),
			BeforeFormat: imgconv.NewImage(from),
			AfterFormat:  imgconv.NewImage(to),
		}

		if !dryRun {
			if err := imgconv.Do(param); err != nil {
				onExit(err)
			}
		} else {
			e := len(param.Path) - len(from)
			fmt.Printf("%s => %s%s\n", path, path[:e], to)
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

func isDir(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fi.IsDir(), nil
}

func getTargetFilepaths(ds []string, ext string) ([]string, error) {
	var names []string

	for _, n := range ds {
		if err := filepath.Walk(n, func(name string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if filepath.Ext(name) == "."+ext {
				names = append(names, name)
			}

			return nil
		}); err != nil {
			return nil, err
		}
	}

	return names, nil
}
