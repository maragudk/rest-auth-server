package handlers

import (
	"context"
	"encoding/gob"
	"encoding/json"
	"net/http"

	"github.com/maragudk/rest-auth-server/model"
)

const (
	sessionUserKey = "user"
)

type sessionGetter interface {
	Exists(ctx context.Context, key string) bool
	Get(ctx context.Context, key string) interface{}
}

type sessionPutter interface {
	RenewToken(ctx context.Context) error
	Put(ctx context.Context, key string, value interface{})
}

type sessionDestroyer interface {
	Destroy(ctx context.Context) error
}

func getUserFromContext(r *http.Request) model.User {
	return r.Context().Value(sessionUserKey).(model.User)
}

type signupper interface {
	Signup(name, password string) error
}

func SignupHandler(repo signupper) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		name := r.Form.Get("name")
		password := r.Form.Get("password")
		if name == "" || password == "" {
			http.Error(w, "name and/or password empty", http.StatusBadRequest)
			return
		}

		if err := repo.Signup(name, password); err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
	}
}

type loginner interface {
	Login(name, password string) (*model.User, error)
}

func LoginHandler(repo loginner, s sessionPutter) http.HandlerFunc {
	// Register our user type with the gob encoding used by the session handler
	gob.Register(model.User{})

	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		name := r.Form.Get("name")
		password := r.Form.Get("password")

		if name == "" || password == "" {
			http.Error(w, "name and/or password empty", http.StatusBadRequest)
			return
		}

		user, err := repo.Login(name, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		if user == nil {
			http.Error(w, "email and/or password incorrect", http.StatusForbidden)
			return
		}

		// Renew the token to avoid session fixation attacks
		if err := s.RenewToken(r.Context()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Put the whole user info into the session
		s.Put(r.Context(), sessionUserKey, user)
	}
}

func LogoutHandler(s sessionDestroyer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := s.Destroy(r.Context()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// CheckSessionHandler sits behind the auth middleware and just returns the current user info.
func CheckSessionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := getUserFromContext(r)
		userAsJSON, err := json.Marshal(user)
		if err != nil {
			panic(err)
		}
		_, _ = w.Write(userAsJSON)
	}
}
