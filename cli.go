package main

import (
	"flag"
)

var serverFlag bool

func init() {
	flag.BoolVar(&serverFlag, "server", false, "is enable server")
	flag.Parse()
}

func main() {
	if serverFlag == true {
		var httpServer = NewHttpServer("0.0.0.0:8081")
		go httpServer.ListenAndServe()

		var server = NewTcpServer(":2000")
		server.ListenAndServe(conn_handler)
	} else {
		var client = NewTcpClient(":2000")
		client.Start()
	}
}
