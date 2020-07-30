/*
converter は画像変換をするパッケージになります。
*/
package converter

import (
	"errors"
	"fmt"
	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// Converter は変換元・変換後のディレクトリを表しますが、変換に関するメソッドも持っています。
type Converter struct {
	// 変換元ディレクトリ
	fromDir string
	// 変換後ディレクトリ
	toDir string
}

// SetDir は Convertメソッドで画像変換する変換元ディレクトリと変換先ディレクトリを指定します。
func (c *Converter) SetDir(fromDir string, toDir string) error {
	// 変換元画像ディレクトリの存在するか確認する。
	f, err := os.Stat(fromDir)
	if os.IsNotExist(err) || !f.IsDir() {
		return errors.New("ディレクトリが存在しないか、ディレクトリではありません。")
	}
	c.fromDir = fromDir

	if len(toDir) == 0 {
		c.toDir = fromDir
	} else {
		c.toDir = toDir
	}

	return nil
}

//　Convert は 指定された変換元画像フォーマットを同じく指定された変換先画像フォーマットに変換します。
//　変換元ディレクトリと変換先ディレクトリは SetDirメソッドで指定し、変換元ディレクトリ内にあるディレクトリを
//　再帰的に検索して画像変換します。
func (c *Converter) Convert(fromType string, toType string) error {

	fType := checkType(fromType)
	if fType == "unknown" {
		return errors.New("指定された変換する画像フォーマットは対応していません。")
	}

	tType := checkType(toType)
	if tType == "unknown" {
		return errors.New("指定された変換後の画像フォーマットは対応していません。")
	}

	return filepath.Walk(c.fromDir,
		func(path string, info os.FileInfo, err error) error {
			// 変換元ディレクトリもpathに入るので、読み飛ばす。
			if path == c.fromDir {
				return nil
			}

			if info.IsDir() {
				// ディレクトリの場合は読み飛ばす。
				return nil
			} else {
				iFile, _ := os.Open(path)
				defer iFile.Close()

				img, format, _ := image.Decode(iFile)

				// イメージファイルのフォーマットが指定変換元フォーマットか確認
				if fType != format {
					return nil
				}

				var err error
				var out *os.File

				switch tType {
				case "gif":
					out, err = makeOutFile(path, "gif", c.fromDir, c.toDir)
					if err == nil {
						err = gif.Encode(out, img, nil)
					}
				case "jpeg":
					out, err = makeOutFile(path, "jpg", c.fromDir, c.toDir)
					if err == nil {
						err = jpeg.Encode(out, img, nil)
					}
				case "png":
					out, err = makeOutFile(path, "png", c.fromDir, c.toDir)
					if err == nil {
						err = png.Encode(out, img)
					}
				case "bmp":
					out, err = makeOutFile(path, "bmp", c.fromDir, c.toDir)
					if err == nil {
						err = bmp.Encode(out, img)
					}
				case "tiff":
					out, err = makeOutFile(path, "tiff", c.fromDir, c.toDir)
					if err == nil {
						err = tiff.Encode(out, img, nil)
					}
				}
				return err
			}
		})
}

// checkType は指定された画像フォーマットが対応しているか確認します。
func checkType(imageType string) string {
	switch imageType {
	case "gif":
		return "gif"
	case "jpeg", "jpg":
		return "jpeg"
	case "png":
		return "png"
	case "bmp":
		return "bmp"
	case "tiff":
		return "tiff"
	default:
		return "unknown"
	}
}

// makeOutFile は指定されたファイルパス、変換後拡張子、ベースディレクトリから出力先ディレクトリ下に変換後の画像ディレクトリとファイルを作成します。
func makeOutFile(filePath string, toExt string, baseDir string, toDir string) (*os.File, error) {
	// ディレクトリに関する処理
	dir := filepath.Dir(filepath.Clean(filePath))
	toDir = filepath.Join(toDir, strings.Replace(dir, baseDir, "", -1))
	_, err := os.Stat(toDir)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(toDir, 0755); err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	// ファイルに関する処理
	ext := filepath.Ext(filePath)
	fileName := filepath.Base(filePath)
	outFilePath := filepath.Join(toDir, fileName[:len(fileName)-len(ext)+1]+toExt)
	_, err = os.Stat(outFilePath)
	if os.IsNotExist(err) {
		f, err := os.Create(outFilePath)
		return f, err
	} else if err != nil {
		return nil, err
	}
	return nil, fmt.Errorf("変換先ファイルが存在しているので変換を行いません。:%s", filePath)
}
