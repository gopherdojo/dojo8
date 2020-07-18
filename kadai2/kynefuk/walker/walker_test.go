package walker

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
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

func CreateTestDir() string {
	dir, err := ioutil.TempDir("./", "example")
	if err != nil {
		log.Fatalf("failed to create test dir. error: %s", err)
	}
	return dir
}

func CreateTmpFile(dir string) []string {
	var tmpFilePaths []string
	for _, v := range fileExtList {
		f, err := ioutil.TempFile(dir, "example.*."+v)
		if err != nil {
			log.Fatalf("failed to create tmp file. error: %s", err)
		}
		tmpFilePaths = append(tmpFilePaths, f.Name())
	}
	return tmpFilePaths
}

func TestWalker(t *testing.T) {
	testDir := CreateTestDir()
	tmpFilePaths := CreateTmpFile(testDir)
	defer os.RemoveAll(testDir)

	type TestCase struct {
		name, directory, fromFmt string
		files                    []string
		err                      error
	}

	type TestCases []struct {
		name, directory, fromFmt string
		files                    []string
		err                      error
	}

	var testCases TestCases
	for _, fileFmt := range fileExtList {
		name := fmt.Sprintf("walker collects %s file", fileFmt)
		tests := TestCase{
			name: name, directory: testDir, fromFmt: fileFmt, files: tmpFilePaths, err: nil,
		}
		testCases = append(testCases, tests)
	}

	for _, tt := range testCases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			dirWalker := NewWalker(tt.fromFmt)
			files, err := dirWalker.Dirwalk(tt.directory)
			if err != tt.err {
				t.Errorf("test failed. error: %s", err)
			}
			if reflect.DeepEqual(tt.files, files) {
				t.Errorf("want %s, but got %s", tt.files, files)
			}
		})
	}
}
