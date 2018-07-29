package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"bufio"
	"os"
)

func write(conn net.Conn) {
	stdReader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := stdReader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatalln(err)
			}
		}

		content := make([]byte, len(line) + 1)
		copy(content, line)
		content[len(line)] = '\t'

		conn.Write([]byte(content))

		fmt.Printf("[Write]: `%s`\n", content)
	}

	conn.Close()
}

func read(conn net.Conn, quit chan<- bool) {
	r := bufio.NewReader(conn)

	for {
		buf, err := r.ReadBytes('\t')
		if err == io.EOF {
			fmt.Printf("Server close the connection\n")
			break
		}
		if err != nil {
			fmt.Println(err)
			log.Fatalln(err)
		}
		//TODO: 如何正确地关闭TCP连接，这里报错:
		// read tcp 127.0.0.1:65438->127.0.0.1:2000: use of closed network connection
		fmt.Printf("[Read]:`%s`\n", buf)

	}

	quit<-true

}

func main() {
	quit := make(chan bool)
	addr := ":2000"

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	go read(conn, quit)
	write(conn)

	<-quit
}
