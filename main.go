package main

import (
	"net/http"
)

func StudentHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to Student page"))
	}
}

func RequestMethodGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte("Method is not allowed"))
			return
		}

		// untuk meneruskan ke Function Handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	// TODO: answer here
	mux := http.DefaultServeMux
	mux.HandleFunc("/student", StudentHandler())

	handler := RequestMethodGet(mux)

	http.ListenAndServe("localhost:8080", handler)
}
