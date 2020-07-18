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
func (cli *Command) Run(args *Args) int {
	dirWalker := walker.NewWalker(args.from)
	files, err := dirWalker.Dirwalk(args.directory)
	if err != nil {
		fmt.Fprintf(cli.ErrStream, "failed to read directory: %s, err: %s\n", args.directory, err)
		return ExitCodeError
	}

	imgConverter := converter.NewConverter(args.from, args.to)
	for _, file := range files {
		if err := imgConverter.ConvertFormat(file); err != nil {
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
