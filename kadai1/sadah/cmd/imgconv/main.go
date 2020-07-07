package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/gopherdojo/dojo8/kadai1/sadah/imgconv"
)

const (
	name          = "imgconv"
	version       = "0.0.1"
	defalutPath   = "./"
	defalutSrcExt = "jpg"
	defalutTarExt = "png"
	usage         = `imgconv is an image converter
Usage: imgconv [options...] [path]
Use "imgconv --help" for more information about a command.

Supported formats:`
	optVerboseText     = "Print verbose messages"
	optShowVersionText = "Show version"
	optSrcExtText      = "Set a source extension"
	optTarExtText      = "Set a target extension"
)

var (
	verbose     bool
	showVersion bool
	srcExt      string
	tarExt      string
	path        string
)

func init() {
	flag.BoolVar(&verbose, "verbose", false, optVerboseText)

	flag.BoolVar(&showVersion, "v", false, optShowVersionText)
	flag.BoolVar(&showVersion, "version", false, optShowVersionText)

	flag.StringVar(&srcExt, "s", defalutSrcExt, optSrcExtText)
	flag.StringVar(&srcExt, "source", defalutSrcExt, optSrcExtText)

	flag.StringVar(&tarExt, "t", defalutTarExt, optTarExtText)
	flag.StringVar(&tarExt, "target", defalutTarExt, optTarExtText)

	flag.Usage = func() {
		usageTxt := usage
		fmt.Fprintf(os.Stderr, "%s %s\n\n", usageTxt, imgconv.SupportedExts)
		flag.PrintDefaults()
	}
	flag.Parse()

	if verbose {
		log.SetOutput(os.Stderr)
	} else {
		log.SetOutput(ioutil.Discard)
	}

	if showVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	arg0 := flag.Arg(0)
	if arg0 == "" {
		path = defalutPath
	} else {
		path = arg0
	}

	srcExt = "." + srcExt
	tarExt = "." + tarExt
}

func run() (err error) {
	log.Println("Target Path: " + path)
	log.Println("Source Extension: " + srcExt)
	log.Println("Target Extension: " + tarExt)

	if err = imgconv.VerifySupportedExt(srcExt); err != nil {
		return
	}

	if err = imgconv.VerifySupportedExt(tarExt); err != nil {
		return
	}

	return filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filepath.Ext(path) == srcExt {
				return imgconv.Conv(path, tarExt)
			}
			return nil
		})
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
