package main

import (
	"fmt"
	"log"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tcpServer := GetTcpServer()
	tcpServer.sessManager.VisitSession(func(sess Session) bool {
		conn := sess.Raw()
		log.Println("Ready to write", conn.RemoteAddr())

		conn.Write([]byte("hello, world!\t"))
		return true
	})

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
