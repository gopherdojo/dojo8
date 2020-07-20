package imgconv

import (
	"io"
	"os"
)

type OpenCreator interface {
	Open(string) (io.ReadCloser, error)
	Create(string) (io.WriteCloser, error)
}

type File struct{}

func (File) Open(n string) (io.ReadCloser, error)    { return os.Open(n) }
func (File) Create(n string) (io.WriteCloser, error) { return os.Create(n) }

func NewFile() File {
	return File{}
}
