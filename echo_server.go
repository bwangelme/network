package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
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
			r := bufio.NewReader(conn)
			var total = 0

			for {
				buf, err := r.ReadBytes('\t')
				if err != nil {
					log.Fatalln(err)
				}

				fmt.Println(string(buf))

				if string(buf) == "exit\t" {
					break
				}

				n, err := conn.Write(buf)
				if err != nil {
					log.Fatalln(err)
				}

				total += n
			}

			fmt.Printf("Total Copy %d bytes\n", total)
			conn.Close()
		}(conn)
	}
}
