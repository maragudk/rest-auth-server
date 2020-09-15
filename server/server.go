package server

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"

	"github.com/maragudk/rest-auth-server/storage"
)

type Server struct {
	address string
	log     *log.Logger
	mux     *chi.Mux
	sm      *scs.SessionManager
	storer  *storage.Storer
}

type Options struct {
	Address string
	Logger  *log.Logger
	Storer  *storage.Storer
}

func New(opts Options) *Server {
	return &Server{
		address: opts.Address,
		log:     opts.Logger,
		mux:     chi.NewMux(),
		storer:  opts.Storer,
	}
}

func (s *Server) Start() error {
	s.sm = scs.New()

	s.setupRoutes()

	server := http.Server{
		Addr:         s.address,
		Handler:      s.mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
		ErrorLog:     s.log,
	}

	s.log.Printf("Listening on https://%v\n", s.address)
	if err := server.ListenAndServeTLS("cert.pem", "key.pem"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("could not listen and serve: %w", err)
	}

	return nil
}
