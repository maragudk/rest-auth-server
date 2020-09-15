package handlers

import (
	"context"
	"net/http"

	"github.com/maragudk/rest-auth-server/model"
)

type Middleware = func(http.Handler) http.Handler

// Authorize creates Middleware that checks that there's a user logged in.
func Authorize(s sessionGetter) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !s.Exists(r.Context(), sessionUserKey) {
				http.Error(w, "unauthorized, please login", http.StatusUnauthorized)
				return
			}

			user, ok := s.Get(r.Context(), sessionUserKey).(model.User)
			if !ok {
				http.Error(w, "could not hydrate user", http.StatusInternalServerError)
				return
			}

			// We store the user directly in the context instead of having to use the session manager
			ctx := context.WithValue(r.Context(), sessionUserKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
