package imgconv_test

import (
	"image"
	"image/color"
	"io"
)

type mockCloser struct{}

func (mockCloser) Read(p []byte) (int, error)  { return 0, nil }
func (mockCloser) Close() error                { return nil }
func (mockCloser) Write(p []byte) (int, error) { return 0, nil }

type openFunc func(string) (io.ReadCloser, error)
type createFunc func(string) (io.WriteCloser, error)

type mockFileHandler struct {
	mockOpen   openFunc
	mockCreate createFunc
}

func (m mockFileHandler) Open(s string) (io.ReadCloser, error)    { return m.mockOpen(s) }
func (m mockFileHandler) Create(s string) (io.WriteCloser, error) { return m.mockCreate(s) }

func newMockFileHandler(of openFunc, cf createFunc) mockFileHandler {
	return mockFileHandler{
		mockOpen:   of,
		mockCreate: cf,
	}
}

type decodeFunc func(io.Reader) (image.Image, error)
type encodeFunc func(io.Writer, image.Image) error
type getExtFunc func() string

type mockImage struct {
	mockDecode decodeFunc
	mockEncode encodeFunc
	mockGetExt getExtFunc
}

func (m mockImage) Decode(r io.Reader) (image.Image, error) { return m.mockDecode(r) }
func (m mockImage) Encode(w io.Writer, i image.Image) error {
	return m.mockEncode(w, i)
}
func (m mockImage) GetExt() string { return m.mockGetExt() }

func newMockImage(df decodeFunc, ef encodeFunc, gf getExtFunc) mockImage {
	return mockImage{
		mockDecode: df,
		mockEncode: ef,
		mockGetExt: gf,
	}
}

type mockStdImage struct{}

func (mockStdImage) ColorModel() (c color.Model) { return }
func (mockStdImage) Bounds() (r image.Rectangle) { return }
func (mockStdImage) At(int, int) (c color.Color) { return }
