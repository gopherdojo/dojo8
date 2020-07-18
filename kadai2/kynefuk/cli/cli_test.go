package cli

import (
	"bytes"
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
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

func createTestData() string {
	// setup
	testDir, err := ioutil.TempDir("", "testdata")
	if err != nil {
		log.Fatal(err)
	}

	testFilePath := "../testdata/gopher.png"

	file, err := os.Open(testFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	for _, ext := range fileExtList {
		out, _ := os.Create(ConvertExt(testFilePath, ext))
		out.Close()

	}
	return testDir
}

func TestCLI(t *testing.T) {
	testDir := createTestData()
	defer os.Remove(testDir)
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	command := NewCommand(outStream, errStream)

	tests := []struct {
		name, directory, fromFormat, toFormat, outStream, errStream string
		exitCode                                                    int
	}{
		{name: "success", directory: testDir, fromFormat: "png", toFormat: "jpg", outStream: "", errStream: "", exitCode: 0},
		{name: "invalid directory", directory: "dummyDir", fromFormat: "jpg", toFormat: "png", outStream: "", errStream: "failed to read directory: dummyDir, err: open dummyDir: no such file or directory", exitCode: 1},
		{name: "invalid to format", directory: testDir, fromFormat: "jpg", toFormat: "hoge", outStream: "", errStream: "failed to convert img, err: unknown format type", exitCode: 1},
		{name: "invalid from format", directory: testDir, fromFormat: "hoge", toFormat: "png", outStream: "", errStream: "", exitCode: 0},
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
