//Package converter は指定されたパスに対し、画像変換を行います。
//（対応形式：jpg,png,gif,bmp,tiff）
package converter

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"regexp"
	"strings"

	"golang.org/x/image/bmp"
	"golang.org/x/image/tiff"
)

//Args ImageConverter()を使う際に指定が必要となる引数
type Args struct {
	FilePath     string //変換対象となるファイル名
	BeforeFormat string //変換前　画像形式
	AfterFormat  string //変換後　画像形式
}

//ImageConverter 指定した画像形式に変換を行います。
func ImageConverter(a Args) error {

	f, err := os.Open(a.FilePath)
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	defer f.Close()

	//ファイルオブジェクトを画像オブジェクトに変換
	img, _, err := image.Decode(f)
	if err != nil {
		return fmt.Errorf("%w : %s", err, f.Name())
	}

	//変換後のファイルパス
	rep := regexp.MustCompile("(?i)" + a.BeforeFormat + "$")
	outFilePath := rep.ReplaceAllString(f.Name(), a.AfterFormat)

	//変換後ファイルの新規作成
	out, err := os.Create(outFilePath)
	if err != nil {
		return fmt.Errorf("%w : %s", err, f.Name())
	}

	//変換する画像形式に応じてエンコードする
	switch strings.ToLower(a.AfterFormat) {
	case "jpg":
		err = jpeg.Encode(out, img, &jpeg.Options{Quality: jpeg.DefaultQuality})
	case "png":
		err = png.Encode(out, img)
	case "gif":
		err = gif.Encode(out, img, &gif.Options{NumColors: 256, Quantizer: nil, Drawer: nil})
	case "bmp":
		err = bmp.Encode(out, img)
	case "tiff":
		err = tiff.Encode(out, img, &tiff.Options{Compression: 0, Predictor: true})
	default:
		return fmt.Errorf("The specified image format is not supported. : " + f.Name())
	}

	//エンコード時にエラーが返ってきていないかチェック
	if err != nil {
		return fmt.Errorf("%w : %s", err, out.Name())
	}

	if err := out.Close(); err != nil {
		return fmt.Errorf("%w : %s", err, out.Name())
	}

	return nil

}
