package walker

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

// Walker walk directory to extract files
type Walker struct {
	TargetExt string
}

// Dirwalk walk directory and collect files
func (walker *Walker) Dirwalk(directory string) ([]string, error) {
	items, err := ioutil.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	var paths []string
	for _, item := range items {
		if item.IsDir() {
			items, err := walker.Dirwalk(filepath.Join(directory, item.Name()))
			if err != nil {
				return nil, err
			}
			paths = append(paths, items...)
			continue
		}
		if strings.HasSuffix(item.Name(), walker.TargetExt) {
			paths = append(paths, filepath.Join(directory, item.Name()))
		}
	}
	return paths, nil
}

// NewWalker is a constructor of Walker
func NewWalker(targetExt string) *Walker {
	return &Walker{TargetExt: targetExt}
}
