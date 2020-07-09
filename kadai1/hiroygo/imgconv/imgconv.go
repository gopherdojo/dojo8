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
	ext := strings.ToLower(filepath.Ext(name))
	ext = strings.Replace(ext, ".", "", -1)
	return ToImageType(ext)
}

// GetImageFilePathesRecursive ディレクトリ内の指定された画像種別のファイルパスを再帰的に返す
func GetImageFilePathesRecursive(dirPath string, target ImageType) ([]string, error) {
	var pathes []string
	err := filepath.Walk(dirPath, func(visitPath string, f os.FileInfo, err error) error {
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

// SaveImage Image を指定された画像形式で保存する
func SaveImage(m image.Image, out ImageType, savePath string) (rerr error) {
	if m == nil {
		rerr = fmt.Errorf("Image が nil です")
		return
	}

	f, err := os.Create(savePath)
	if err != nil {
		rerr = fmt.Errorf("Create error, %w", err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			rerr = fmt.Errorf("Close error, %w", err)
		}
	}()

	switch out {
	case Jpeg:
		err = jpeg.Encode(f, m, nil)
	case Png:
		err = png.Encode(f, m)
	case Tiff:
		err = tiff.Encode(f, m, nil)
	case Bmp:
		err = bmp.Encode(f, m)
	case Gif:
		err = gif.Encode(f, m, nil)
	default:
		err = fmt.Errorf("画像種別 %s は不明です", out.String())
	}
	if err != nil {
		rerr = fmt.Errorf("Encode error, %w", err)
	}

	return
}
