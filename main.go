package main

import (
	"fmt"
	"net/http"
)

func main() {
	sm := http.NewServeMux()
	s := http.Server{
		Addr:    ":8080",
		Handler: sm,
	}
	s.ListenAndServe()
	fmt.Println("Hello, chirpy")
}
