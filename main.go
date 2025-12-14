package main

import (
	"net/http"
)

type HealthHandler struct{}

func (h *HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("OK"))
}

func main() {
	mux := http.NewServeMux()
	healthHandler := HealthHandler{}
	mux.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir("."))))
	mux.Handle("/app/assets", http.StripPrefix("/app/", http.FileServer(http.Dir("./"))))
	mux.HandleFunc("/healthz", healthHandler.ServeHTTP)
	s := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	s.ListenAndServe()
}
