package downloader

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_downloader_ContentLength(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		contentLength  string
		httpStatusCode int
		want           int64
		wantErr        bool
	}{
		{name: "contentLengthがなければ-1が返る", contentLength: "", httpStatusCode: http.StatusOK, want: -1},
		{name: "contentLengthが0なら0が返る", contentLength: "0", httpStatusCode: http.StatusOK, want: 0},
		{name: "contentLengthが10なら10が返る", contentLength: "10", httpStatusCode: http.StatusOK, want: 10},
		{name: "httpStatusが200以外でもcontentLengthはそのまま返される", contentLength: "10", httpStatusCode: http.StatusNotFound, want: 10},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if tt.contentLength != "" {
					w.Header().Add("Content-Length", tt.contentLength)
				}
				w.WriteHeader(tt.httpStatusCode)
				_, _ = w.Write([]byte{})
			}))

			d := &downloader{url: ts.URL}
			got, err := d.ContentLength()
			if (err != nil) != tt.wantErr {
				t.Errorf("ContentLength() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ContentLength() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_downloader_AcceptRanges(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name           string
		acceptRanges   string
		httpStatusCode int
		want           string
		wantErr        bool
	}{
		{name: "acceptRangesがなければ-1が返る", acceptRanges: "", httpStatusCode: http.StatusOK, want: "none"},
		{name: "acceptRangesが0なら0が返る", acceptRanges: "none", httpStatusCode: http.StatusOK, want: "none"},
		{name: "acceptRangesが10なら10が返る", acceptRanges: "byte", httpStatusCode: http.StatusOK, want: "byte"},
		{name: "httpStatusが200以外でもacceptRangesはそのまま返される", acceptRanges: "byte", httpStatusCode: http.StatusNotFound, want: "byte"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if tt.acceptRanges != "" {
					w.Header().Add("Accept-Ranges", tt.acceptRanges)
				}
				w.WriteHeader(tt.httpStatusCode)
				_, _ = w.Write([]byte{})
			}))
			d := &downloader{url: ts.URL}
			got, err := d.AcceptRanges()
			if (err != nil) != tt.wantErr {
				t.Errorf("AcceptRanges() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AcceptRanges() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	t.Parallel()
	url := "http://example.com/foo/bar"
	want := &downloader{url: url}
	got := New(url)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("New() = %v, want %v", got, want)
	}
}
