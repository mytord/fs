package middlewares

import (
	"context"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userIdHeader := r.Header.Get("X-User-Id")

		if userIdHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userId", userIdHeader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
