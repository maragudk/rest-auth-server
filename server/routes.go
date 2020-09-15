package server

import (
	"github.com/go-chi/chi"
	"github.com/maragudk/rest-auth-server/handlers"
)

func (s *Server) setupRoutes() {
	s.mux.Use(s.sm.LoadAndSave)

	s.mux.Post("/login", handlers.LoginHandler(s.storer, s.sm))
	s.mux.Post("/logout", handlers.LogoutHandler(s.sm))
	s.mux.Post("/signup", handlers.SignupHandler(s.storer))

	s.mux.Group(func(r chi.Router) {
		r.Use(handlers.Authorize(s.sm))

		r.Get("/check", handlers.CheckSessionHandler())
	})
}
