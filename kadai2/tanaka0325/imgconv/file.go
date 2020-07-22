package imgconv

import (
	"io"
	"os"
)

type FileHandler interface {
	Open(string) (io.ReadCloser, error)
	Create(string) (io.WriteCloser, error)
}

type File struct {
	Reader io.Reader
	Writer io.Writer
}

func (File) Open(n string) (io.ReadCloser, error)    { return os.Open(n) }
func (File) Create(n string) (io.WriteCloser, error) { return os.Create(n) }

func NewFile() File {
	return File{}
}
