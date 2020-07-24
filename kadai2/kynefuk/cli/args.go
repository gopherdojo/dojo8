package cli

import (
	"fmt"
)

// Args represents CLI's arguments object
type Args struct {
	DirecTory, From, To string
}

// Validate is a validation method
func (args *Args) Validate() error {
	fileExtMap := createFileExtMap()
	if _, ok := fileExtMap[args.From]; !ok {
		return fmt.Errorf("argument of \"-f, --From\" is not valid file format. invalid format: %s", args.From)
	}
	if _, ok := fileExtMap[args.To]; !ok {
		return fmt.Errorf("argument of \"-t, --To\" is not valid file format. invalid format: %s", args.To)
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

// NewArgs is a construcTor of Args
func NewArgs(DirecTory, From, To string) *Args {
	return &Args{DirecTory: DirecTory, From: From, To: To}
}
