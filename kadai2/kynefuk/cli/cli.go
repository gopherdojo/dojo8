package cli

import (
	"fmt"
	"io"

	"github.com/gopherdojo/dojo8/kadai2/kynefuk/converter"
	"github.com/gopherdojo/dojo8/kadai2/kynefuk/walker"
)

// it represents Exit Code
const (
	ExitCodeOK = iota
	ExitCodeError
)

// Command represents CLI object
type Command struct {
	OutStream, ErrStream io.Writer
}

// Run invoke cli main logic
func (cli *Command) Run(directory, from, to string) int {
	dirWalker := walker.NewWalker(from)
	files, err := dirWalker.Dirwalk(directory)
	if err != nil {
		fmt.Fprintf(cli.ErrStream, "failed to read directory: %s, err: %s\n", directory, err)
		return ExitCodeError
	}

	imgConverter := converter.NewConverter(from, to)
	for _, file := range files {
		dstPath := converter.ConvertExt(file, from, to)
		if err := imgConverter.ConvertFormat(file, dstPath); err != nil {
			fmt.Fprintf(cli.ErrStream, "failed to convert img, err: %s\n", err)
			return ExitCodeError
		}
	}
	return ExitCodeOK
}

// NewCommand is a constructor of CLI
func NewCommand(outStream, errStream io.Writer) *Command {
	return &Command{OutStream: outStream, ErrStream: errStream}
}
