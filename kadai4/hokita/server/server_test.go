package server_test

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gopherdojo/dojo8/kadai4/hokita/omikuji/omikuji"
	"github.com/gopherdojo/dojo8/kadai4/hokita/omikuji/server"
)

func TestGETOmikuji(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		time time.Time
		want omikuji.Result
	}{
		"standard day": {
			time: time.Date(2000, 1, 4, 0, 0, 0, 0, time.Local),
			want: omikuji.Result{
				Result: "大凶",
			},
		},
		"new year": {
			time: time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local),
			want: omikuji.Result{
				Result: "大吉",
			},
		},
	}

	for name, test := range tests {
		rand.Seed(1)

		server := &server.Server{
			omikuji.New(test.time),
		}

		request, _ := http.NewRequest(http.MethodGet, "/omikuji/", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		t.Run(name, func(t *testing.T) {
			var got omikuji.Result

			err := json.NewDecoder(response.Body).Decode(&got)
			if err != nil {
				t.Fatalf(`unable to parse json. error: "%v"`, got)
			}
			if got != test.want {
				t.Fatalf(`want:"%v" actual:"%v"`, test.want, got)
			}
		})
	}
}
