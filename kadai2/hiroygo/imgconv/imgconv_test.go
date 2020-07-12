package imgconv_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/gopherdojo/dojo8/kadai2/hiroygo/imgconv"
)

var testdataDir string = filepath.Join("..", "testdata")
var testdataSubDir string = filepath.Join(testdataDir, "sub")
var testJpgFile string = filepath.Join(testdataDir, "go.jpg")
var testJpegFile string = filepath.Join(testdataDir, "go.jpeg")
var testSubJpgFile string = filepath.Join(testdataSubDir, "go.jpg")
var testPngFile string = filepath.Join(testdataDir, "go.png")
var testTiffFile string = filepath.Join(testdataDir, "go.tiff")
var testBmpFile string = filepath.Join(testdataDir, "go.bmp")
var testGifFile string = filepath.Join(testdataDir, "go.gif")

func TestToImageType(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		expected imgconv.ImageType
	}{
		{name: "jpg", input: "jpg", expected: imgconv.Jpeg},
		{name: "jpeg", input: "jpeg", expected: imgconv.Jpeg},
		{name: "png", input: "png", expected: imgconv.Png},
		{name: "tiff", input: "tiff", expected: imgconv.Tiff},
		{name: "bmp", input: "bmp", expected: imgconv.Bmp},
		{name: "gif", input: "gif", expected: imgconv.Gif},
		{name: "txt", input: "txt", expected: imgconv.Unknown},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if actual := imgconv.ToImageType(c.input); actual != c.expected {
				t.Errorf("want ToImageType(%s) = %v, got %v", c.input, c.expected, actual)
			}
		})
	}
}

func TestImageFilePathesRecursive(t *testing.T) {
	cases := []struct {
		name           string
		dir            string
		imgType        imgconv.ImageType
		expectedPathes []string
	}{
		{name: "jpg_files", dir: testdataDir, imgType: imgconv.Jpeg, expectedPathes: []string{testJpegFile, testJpgFile, testSubJpgFile}},
		{name: "png_files", dir: testdataDir, imgType: imgconv.Png, expectedPathes: []string{testPngFile}},
		{name: "tiff_files", dir: testdataDir, imgType: imgconv.Tiff, expectedPathes: []string{testTiffFile}},
		{name: "bmp_files", dir: testdataDir, imgType: imgconv.Bmp, expectedPathes: []string{testBmpFile}},
		{name: "gif_files", dir: testdataDir, imgType: imgconv.Gif, expectedPathes: []string{testGifFile}},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual, err := imgconv.ImageFilePathesRecursive(c.dir, c.imgType)
			if err != nil {
				t.Fatalf("ImageFilePathesRecursive(%s, %v) error %v", c.dir, c.imgType, err)
			}

			// NOTE:expectedPathes と actual の中身の順序まで一致しないと成功にならない
			// e.g. expectedPathes:{"test1.jpg", "test2.jpg"}, actual:{"test1.jpg", "test2.jpg"} => OK
			// e.g. expectedPathes:{"test1.jpg", "test2.jpg"}, actual:{"test2.jpg", "test1.jpg"} => NG
			if !reflect.DeepEqual(actual, c.expectedPathes) {
				t.Errorf("want ImageFilePathesRecursive(%s, %v) = %v, got %v", c.dir, c.imgType, c.expectedPathes, actual)
			}
		})
	}
}

func TestLoadImage(t *testing.T) {
	cases := []struct {
		name    string
		path    string
		wantErr bool
	}{
		{name: "jpg", path: testJpgFile, wantErr: false},
		{name: "png", path: testPngFile, wantErr: false},
		{name: "tiff", path: testTiffFile, wantErr: false},
		{name: "bmp", path: testBmpFile, wantErr: false},
		{name: "gif", path: testGifFile, wantErr: false},
		{name: "directory", path: testdataDir, wantErr: true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := imgconv.LoadImage(c.path)
			if !c.wantErr && err != nil {
				t.Errorf("LoadImage(%s) error %v", c.path, err)
			}
		})
	}
}

func TestReplaceExt(t *testing.T) {
	cases := []struct {
		name     string
		path     string
		outType  imgconv.ImageType
		expected string
	}{
		{name: "without_ext", path: "test", outType: imgconv.Jpeg, expected: "test.jpg"},
		{name: "ends_dot", path: "test.", outType: imgconv.Jpeg, expected: "test.jpg"},
		{name: "begin_dot", path: ".test.", outType: imgconv.Jpeg, expected: ".test.jpg"},
		{name: "only_dot", path: ".", outType: imgconv.Jpeg, expected: "..jpg"},
		{name: "empty", path: "", outType: imgconv.Jpeg, expected: "..jpg"},
		{name: "has_parent_dir", path: testJpgFile, outType: imgconv.Png, expected: testPngFile},
		{name: "jpg", path: testPngFile, outType: imgconv.Jpeg, expected: testJpgFile},
		{name: "png", path: testJpgFile, outType: imgconv.Png, expected: testPngFile},
		{name: "tiff", path: testJpgFile, outType: imgconv.Tiff, expected: testTiffFile},
		{name: "bmp", path: testJpgFile, outType: imgconv.Bmp, expected: testBmpFile},
		{name: "gif", path: testJpgFile, outType: imgconv.Gif, expected: testGifFile},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := imgconv.ReplaceExt(c.path, c.outType)
			if actual != c.expected {
				t.Errorf("want ReplaceExt(%s, %v) = %s, got %s", c.path, c.outType, c.expected, actual)
			}
		})
	}
}

func tempDir(t *testing.T) string {
	t.Helper()
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("TempDir error %v", err)
	}

	t.Cleanup(func() {
		if err := os.RemoveAll(dir); err != nil {
			t.Fatalf("RemoveAll error %v", err)
		}
	})

	return dir
}

func TestSaveImage(t *testing.T) {
	tmpDir := tempDir(t)
	saveFileBase := filepath.Join(tmpDir, "test")

	type saveParam struct {
		name     string
		saveType imgconv.ImageType
		savePath string
		wantErr  bool
	}
	saveParams := []saveParam{
		{name: "jpg", saveType: imgconv.Jpeg, savePath: imgconv.ReplaceExt(saveFileBase, imgconv.Jpeg), wantErr: false},
		{name: "png", saveType: imgconv.Png, savePath: imgconv.ReplaceExt(saveFileBase, imgconv.Png), wantErr: false},
		{name: "tiff", saveType: imgconv.Tiff, savePath: imgconv.ReplaceExt(saveFileBase, imgconv.Tiff), wantErr: false},
		{name: "bmp", saveType: imgconv.Bmp, savePath: imgconv.ReplaceExt(saveFileBase, imgconv.Bmp), wantErr: false},
		{name: "gif", saveType: imgconv.Gif, savePath: imgconv.ReplaceExt(saveFileBase, imgconv.Gif), wantErr: false},
		{name: "unknown", saveType: imgconv.Unknown, savePath: imgconv.ReplaceExt(saveFileBase, imgconv.Unknown), wantErr: true},
	}

	cases := []struct {
		name       string
		srcPath    string
		saveParams []saveParam
	}{
		{name: "jpg_to", srcPath: testJpgFile, saveParams: saveParams},
		{name: "png_to", srcPath: testPngFile, saveParams: saveParams},
		{name: "tiff_to", srcPath: testTiffFile, saveParams: saveParams},
		{name: "bmp_to", srcPath: testBmpFile, saveParams: saveParams},
		{name: "gif_to", srcPath: testGifFile, saveParams: saveParams},
	}

	for _, c := range cases {
		for _, innerCase := range c.saveParams {
			t.Run(c.name+"_"+innerCase.name, func(t *testing.T) {
				m, err := imgconv.LoadImage(c.srcPath)
				if err != nil {
					t.Fatalf("LoadImage(%s) error %v", c.srcPath, err)
				}

				err = imgconv.SaveImage(m, innerCase.saveType, innerCase.savePath)
				if !innerCase.wantErr && err != nil {
					t.Errorf("SaveImage(image.Image, %v, %s) error %v", innerCase.saveType, innerCase.savePath, err)
				}
			})
		}
	}
}
