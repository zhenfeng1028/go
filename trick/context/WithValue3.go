package main

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	traceid := ""
	if m := r.Context().Value("traceid"); m != nil {
		if value, ok := m.(string); ok {
			traceid = value
		}
	}
	w.Header().Add("traceid", traceid)
	w.Write([]byte("Welcome to China"))
}

func traceID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceid := uuid.New().String()
		ctx := context.WithValue(r.Context(), "traceid", traceid)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

func main() {
	welcomeHandler := http.HandlerFunc(welcome)
	http.Handle("/welcome", traceID(welcomeHandler))
	http.ListenAndServe(":9090", nil)
}
