package imgconv

import (
	"io"
	"os"
)

type OpenCreator interface {
	Open() (io.ReadCloser, error)
	Create() (io.WriteCloser, error)
}

type File struct {
	Path string
}

func (f File) Open() (io.ReadCloser, error)    { return os.Open(f.Path) }
func (f File) Create() (io.WriteCloser, error) { return os.Create(f.Path) }

func NewFile(p string) File {
	return File{Path: p}
}
