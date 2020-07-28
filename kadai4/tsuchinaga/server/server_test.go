package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gopherdojo/dojo8/kadai4/tsuchinaga/clock"
)

func Test_server_GetAddr(t *testing.T) {
	type fields struct {
		addr string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "addrが空文字なら空文字を返す", fields: fields{addr: ""}, want: ""},
		{name: "addrにアドレスが入っていたら入っているアドレスを返す", fields: fields{addr: "127.0.0.1:57024"}, want: "127.0.0.1:57024"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{addr: tt.fields.addr}
			if got := s.GetAddr(); got != tt.want {
				t.Errorf("GetAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}

type testClock struct{ now time.Time }

func (t *testClock) Now() time.Time { return t.now }

func Test_server_FortuneHandler(t *testing.T) {
	type fields struct{ clock clock.Clock }
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "2020/01/01は全て大吉",
			fields: fields{clock: &testClock{time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)}},
			want:   `{"no":1,"foredoom":"大吉","wish":"大吉","health":"大吉","business":"大吉","love":"大吉","study":"大吉"}`},
		{name: "2020/01/02は全て大吉",
			fields: fields{clock: &testClock{time.Date(2020, 1, 2, 0, 0, 0, 0, time.Local)}},
			want:   `{"no":2,"foredoom":"大吉","wish":"大吉","health":"大吉","business":"大吉","love":"大吉","study":"大吉"}`},
		{name: "2020/01/03は全て大吉",
			fields: fields{clock: &testClock{time.Date(2020, 1, 2, 0, 0, 0, 0, time.Local)}},
			want:   `{"no":3,"foredoom":"大吉","wish":"大吉","health":"大吉","business":"大吉","love":"大吉","study":"大吉"}`},
		{name: "2020/01/04はランダム",
			fields: fields{clock: &testClock{time.Date(2020, 1, 4, 0, 0, 0, 0, time.Local)}},
			want:   `{"no":30,"foredoom":"中吉","wish":"凶","health":"大吉","business":"大吉","love":"大吉","study":"凶"}`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{clock: tt.fields.clock}
			mux := http.NewServeMux()
			mux.HandleFunc("/", s.FortuneHandler)
			ts := httptest.NewServer(mux)
			defer ts.Close()
			req, _ := http.NewRequest("GET", ts.URL, nil)
			resp, err := new(http.Client).Do(req)
			if err != nil {
				t.Errorf("%s 失敗\nリクエストエラー: %v\n", t.Name(), err)
				return
			}
			defer resp.Body.Close()
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("%s 失敗\n読み込みエラー: %v\n", t.Name(), err)
				return
			}
			log.Printf("%+v\n", string(b))
		})
	}
}

func Test_server_Run(t *testing.T) {
	want := `{"no":1,"foredoom":"大吉","wish":"大吉","health":"大吉","business":"大吉","love":"大吉","study":"大吉"}`
	serv := &server{clock: &testClock{now: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)}}
	go func() { _ = serv.Run() }()
	for {
		time.Sleep(100 * time.Millisecond)
		if serv.GetAddr() != "" {
			break
		}
	}
	resp, err := http.Get(fmt.Sprintf("http://%s/fortune", serv.GetAddr()))
	if err != nil {
		t.Errorf("%s エラー\n%v\n", t.Name(), err)
		return
	}
	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	if got := string(b); got != want {
		t.Errorf("%s 失敗\n期待: %s\n実際: %s\n", t.Name(), want, got)
	}
}
