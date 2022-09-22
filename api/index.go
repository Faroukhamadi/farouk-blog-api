package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Faroukhamadi/farouk-blog-api/ent"
	"github.com/gorilla/mux"
)

type Server struct {
	l      *log.Logger
	router *mux.Router
	orm    *ent.Client
}

func newServer(l *log.Logger, r *mux.Router, orm *ent.Client) *Server {
	return &Server{l, r, orm}
}

func Init() (srv *http.Server) {
	s, err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}

	srv = &http.Server{
		Addr:    ":4000",
		Handler: s,
	}

	return srv
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func run() (s *Server, err error) {
	client, err := open("postgresql://postgres:faroukhamadi@127.0.0.1/farouk-blog")
	if err != nil {
		return
	}
	// This is mostly useless
	defer client.Close()
	l := log.New(os.Stdout, "farouk-blog-api", log.LstdFlags)
	sm := mux.NewRouter()

	s = newServer(l, sm, client)
	s.routes()

	return
}
