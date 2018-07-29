package main

import (
	"net"
	"log"
	"fmt"
	"io"
)

func main() {
	addr := ":2000"

	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	defer l.Close()
	fmt.Printf("Listening on `%s`\n", addr)

	for {
		conn, err := l.Accept()
		fmt.Printf("Accept Conn from %s\n", conn.RemoteAddr())
		if err != nil {
			log.Fatalln(err)
		}

		go func(conn net.Conn) {
			n, err := io.Copy(conn, conn)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("Total Copy %d bytes\n", n)
			conn.Close()
		}(conn)
	}
}
