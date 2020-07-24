package converter

import (
	"os"

	"github.com/gopherdojo/dojo8/kadai2/kynefuk/helper"

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

const testDirPath = "../testdata"

func TestConverter(t *testing.T) {
	tmpDir := helper.CreateTmpDir()
	defer os.RemoveAll(tmpDir)

	// from→toに画像変換できているかテスト
	for _, from := range helper.FileExtList {
		for _, to := range helper.FileExtList {
			tmpFile := helper.CreateTmpFile(tmpDir, from)
			defer os.Remove(tmpFile.Name())
			imgConverter := NewConverter(from, to)
			err := imgConverter.ConvertFormat(tmpFile.Name(), ConvertExt(tmpFile.Name(), from, to))
			if err != nil {
				t.Errorf("failed to test. error: %s", err)
			}
			exists := helper.Exists(ConvertExt(tmpFile.Name(), from, to))
			if !exists {
				t.Errorf("want true, but got %v", exists)
			}
		}
	}
}
