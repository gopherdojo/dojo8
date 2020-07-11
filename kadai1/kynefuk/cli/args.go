package cli

import (
	"fmt"
)

// Args represents CLI's arguments object
type Args struct {
	directory, from, to string
}

// Validate is a validation method
func (args *Args) Validate() error {
	fileExtMap := createFileExtMap()
	if _, ok := fileExtMap[args.from]; !ok {
		return fmt.Errorf("argument of \"-f, --from\" is not valid file format. invalid format: %s", args.from)
	}
	if _, ok := fileExtMap[args.to]; !ok {
		return fmt.Errorf("argument of \"-t, --to\" is not valid file format. invalid format: %s", args.to)
	}

	return nil
}

func createFileExtMap() map[string]string {
	fileExtMap := make(map[string]string)
	list := []string{
		"jpg",
		"jpeg",
		"png",
		"gif",
		"bmp",
		"tiff",
	}
	for _, v := range list {
		fileExtMap[v] = ""
	}
	return fileExtMap
}

// NewArgs is a constructor of Args
func NewArgs(directory, from, to string) *Args {
	return &Args{directory: directory, from: from, to: to}
}
