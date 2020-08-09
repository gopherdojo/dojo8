package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gopherdojo/dojo8/kadai4/hokita/omikuji/omikuji"
)

type Omikuji interface {
	Draw() *omikuji.Result
}

type Server struct {
	Omikuji Omikuji
}

func Run() {
	os.Exit(run())
}

func run() int {
	server := &Server{NewOmikuji()}

	if err := http.ListenAndServe(":8080", server); err != nil {
		fmt.Fprintf(os.Stderr, "could not listen on port 8080 %v\n", err)
		return 1
	}
	return 0
}

func NewOmikuji() Omikuji {
	return omikuji.New(time.Now())
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := http.NewServeMux()
	router.Handle("/omikuji/", http.HandlerFunc(s.omikujiHandler))

	router.ServeHTTP(w, r)
}

func (s *Server) omikujiHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(s.Omikuji.Draw()); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
