package imgconv

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

const (
	// Unknown 不明な画像
	Unknown ImageType = "unknown"
	// Jpeg Jpeg 画像
	Jpeg ImageType = "jpg"
	// Png Png 画像
	Png ImageType = "png"
	// Tiff Tiff 画像
	Tiff ImageType = "tiff"
	// Bmp Bmp 画像
	Bmp ImageType = "bmp"
	// Gif Gif 画像
	Gif ImageType = "gif"
)

// ImageType 画像種別
type ImageType string

func (t ImageType) String() string {
	return string(t)
}

// ToImageType 文字列を ImageType に変換する
// e.g. "jpg" => Jpeg
func ToImageType(s string) ImageType {
	switch s {
	case "jpeg", "jpg":
		return Jpeg
	case "png":
		return Png
	case "tiff":
		return Tiff
	case "bmp":
		return Bmp
	case "gif":
		return Gif
	default:
		return Unknown
	}
}

func imageTypeFromFileName(name string) ImageType {
	extWithoutDot := strings.TrimLeft(filepath.Ext(name), ".")
	return ToImageType(strings.ToLower(extWithoutDot))
}

// ImageFilePathesRecursive ディレクトリ内の指定された画像種別のファイルパスを再帰的に返す
func ImageFilePathesRecursive(dir string, target ImageType) ([]string, error) {
	var pathes []string
	err := filepath.Walk(dir, func(visitPath string, f os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("WalkFunc error, %w", err)
		}

		if !f.Mode().IsRegular() {
			return nil
		}

		t := imageTypeFromFileName(f.Name())
		if t != target {
			return nil
		}

		pathes = append(pathes, visitPath)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("Walk error, %w", err)
	}

	return pathes, nil
}

// LoadImage 画像を読み込み、その Image を返す
func LoadImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Open error, %w", err)
	}
	defer f.Close()

	m, _, err := image.Decode(f)
	if err != nil {
		return nil, fmt.Errorf("Decode error, %w", err)
	}

	return m, nil
}

// ReplaceExt 画像ファイルパスの拡張子を置換した結果を返す
func ReplaceExt(path string, replace ImageType) string {
	newName := filepath.Base(path[:len(path)-len(filepath.Ext(path))])
	newName += "." + replace.String()
	return filepath.Join(filepath.Dir(path), newName)
}

// SaveImage Image を指定された画像形式で保存する
func SaveImage(m image.Image, out ImageType, savePath string) (rerr error) {
	if m == nil {
		return fmt.Errorf("Image が nil です")
	}

	var encoder func(*os.File, image.Image) error
	switch out {
	case Jpeg:
		encoder = func(f *os.File, m image.Image) error {
			return jpeg.Encode(f, m, nil)
		}
	case Png:
		encoder = func(f *os.File, m image.Image) error {
			return png.Encode(f, m)
		}
	case Tiff:
		encoder = func(f *os.File, m image.Image) error {
			return tiff.Encode(f, m, nil)
		}
	case Bmp:
		encoder = func(f *os.File, m image.Image) error {
			return bmp.Encode(f, m)
		}
	case Gif:
		encoder = func(f *os.File, m image.Image) error {
			return gif.Encode(f, m, nil)
		}
	default:
		return fmt.Errorf("画像種別 %s は不明です", out.String())
	}

	f, err := os.Create(savePath)
	if err != nil {
		return fmt.Errorf("Create error, %w", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			rerr = fmt.Errorf("Close error, %w", err)
		}
	}()

	if err := encoder(f, m); err != nil {
		return fmt.Errorf("Encode error, %w", err)
	}

	return
}
