package validation

// NewValidator - 新しくValidatorを生成する
func NewValidator() Validator {
	return &validator{}
}

// Validator - validationをまとめている構造体のインターフェース
type Validator interface {
	IsValidDir(dir string) bool
	IsValidSrc(src string) bool
	IsValidDest(dest, src string) bool
}

// validator - validationをまとめている構造体
type validator struct{}

// IsValidDir - 許可されたdirかのチェック
func (v validator) IsValidDir(dir string) bool {
	return dir != ""
}

var validFileTypes = map[string]bool{"jpeg": true, "png": true}

// IsValidFileType - 指定されたファイルタイプが利用可能かを返す
func (v validator) IsValidFileType(fileType string) bool {
	return validFileTypes[fileType]
}

// IsValidSrc - 許可されたsrcかのチェック
func (v validator) IsValidSrc(src string) bool {
	return v.IsValidFileType(src)
}

// IsValidDest - 許可されたdestかのチェック
func (v validator) IsValidDest(dest, src string) bool {
	return v.IsValidFileType(dest) && dest != src
}
