package walker

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/gopherdojo/dojo8/kadai2/kynefuk/helper"
)

func TestWalker(t *testing.T) {
	tmpDir := helper.CreateTmpDir()
	tmpFilePaths := helper.CreateTmpFiles(tmpDir)
	defer os.RemoveAll(tmpDir)

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
	for _, fileFmt := range helper.FileExtList {
		name := fmt.Sprintf("walker collects %s file", fileFmt)
		tests := TestCase{
			name: name, directory: tmpDir, fromFmt: fileFmt, files: tmpFilePaths, err: nil,
		}
		testCases = append(testCases, tests)
	}

	for _, c := range testCases {
		t.Run(c.name, func(t *testing.T) {
			dirWalker := NewWalker(c.fromFmt)
			files, err := dirWalker.Dirwalk(c.directory)
			if err != c.err {
				t.Errorf("test failed. error: %s", err)
			}
			if reflect.DeepEqual(c.files, files) {
				helper.ErrorHelper(t, c.files, files)
			}
		})
	}
}
