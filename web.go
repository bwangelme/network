package main

import (
	"fmt"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello, world!")
}

func NewHttpServer(addr string) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/", HomeHandler)

	var srv = &http.Server{
		Handler: mux,
		Addr:    addr,
	}

	return srv
}
