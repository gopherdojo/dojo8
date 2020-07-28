package server

import (
	"context"
	"encoding/json"
	"net"
	"net/http"

	"github.com/gopherdojo/dojo8/kadai4/tsuchinaga/clock"

	"github.com/gopherdojo/dojo8/kadai4/tsuchinaga/fortune"
)

func New(clock clock.Clock) Server {
	return &server{clock: clock}
}

type Server interface {
	GetAddr() string
	Run() error
	Stop(ctx context.Context) error
}

type server struct {
	addr  string
	serv  *http.Server
	clock clock.Clock
}

func (s *server) GetAddr() string {
	return s.addr
}

func (s *server) Run() error {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	defer ln.Close()
	s.addr = ln.Addr().String()

	mux := http.NewServeMux()
	mux.HandleFunc("/fortune", s.FortuneHandler)
	s.serv = &http.Server{Handler: mux}
	return s.serv.Serve(ln)
}

func (s *server) Stop(ctx context.Context) error {
	if s.serv != nil {
		return s.serv.Shutdown(ctx)
	}
	return nil
}

func (s *server) FortuneHandler(w http.ResponseWriter, _ *http.Request) {
	res, err := json.Marshal(fortune.Get(s.clock.Now()))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(res)
}
