package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	command := NewCommand(outStream, errStream)

	tests := []struct {
		name, directory, fromFormat, toFormat, outStream, errStream string
		exitCode                                                    int
	}{
		{name: "success", directory: "../testdata", fromFormat: "png", toFormat: "jpg", outStream: "", errStream: "", exitCode: 0},
		{name: "invalid directory", directory: "dummyDir", fromFormat: "jpg", toFormat: "png", outStream: "", errStream: "failed to read directory: dummyDir, err: open dummyDir: no such file or directory", exitCode: 1},
		{name: "invalid to format", directory: "../testdata", fromFormat: "jpg", toFormat: "hoge", outStream: "", errStream: "", exitCode: 1},
		{name: "invalid from format", directory: "../testdata", fromFormat: "hoge", toFormat: "png", outStream: "", errStream: "", exitCode: 1},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			args := NewArgs(tt.directory, tt.fromFormat, tt.toFormat)
			status := command.Run(args)
			if status != tt.exitCode {
				t.Errorf("want %d, got %d", tt.exitCode, status)
			}
			if !strings.Contains(outStream.String(), tt.outStream) {
				t.Errorf("want %s, got %s", tt.outStream, outStream.String())
			}
			if !strings.Contains(errStream.String(), tt.errStream) {
				t.Errorf("want %s, got %s", tt.errStream, errStream.String())
			}
		})
	}
}
