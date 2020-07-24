package cli

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/gopherdojo/dojo8/kadai2/kynefuk/helper"
)

var fileExtList = []string{
	"jpg",
	"jpeg",
	"png",
	"gif",
	"bmp",
	"tiff",
}

func ConvertExt(filepath, to string) string {
	return strings.Replace(filepath, ".png", to, 1)
}

func TestCLI(t *testing.T) {
	testDir := helper.CreateTmpDir()
	defer os.RemoveAll(testDir)
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	command := NewCommand(outStream, errStream)

	cases := []struct {
		name       string
		directory  string
		fromFormat string
		toFormat   string
		outStream  string
		errStream  string
		exitCode   int
	}{
		{name: "success", directory: testDir, fromFormat: "png", toFormat: "jpg", outStream: "", errStream: "", exitCode: 0},
		{name: "invalid directory", directory: "dummyDir", fromFormat: "jpg", toFormat: "png", outStream: "", errStream: "failed to read directory: dummyDir, err: open dummyDir: no such file or directory", exitCode: 1},
		{name: "invalid to format", directory: testDir, fromFormat: "jpg", toFormat: "hoge", outStream: "", errStream: "failed to convert img, err: unknown format type", exitCode: 1},
		{name: "invalid from format", directory: testDir, fromFormat: "hoge", toFormat: "png", outStream: "", errStream: "", exitCode: 1},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			tmpFile := helper.CreateTmpFile(testDir, c.fromFormat)
			defer os.Remove(tmpFile.Name())
			status := command.Run(c.directory, c.fromFormat, c.toFormat)

			if status != c.exitCode {
				helper.ErrorHelper(t, c.exitCode, status)
			}
			if !strings.Contains(outStream.String(), c.outStream) {
				helper.ErrorHelper(t, c.outStream, outStream.String())
			}
			if !strings.Contains(errStream.String(), c.errStream) {
				helper.ErrorHelper(t, c.errStream, errStream.String())
			}
		})
	}
}
