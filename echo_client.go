package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
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

		content := make([]byte, len(line)+1)
		copy(content, line)
		content[len(line)] = '\t'

		conn.Write([]byte(content))

		fmt.Printf("[Write]: `%s`\n", content)
	}

	conn.Write([]byte("exit\t"))
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
		// Write 已经把 conn 給關閉了，所以read這裏不能夠再使用了
		fmt.Printf("[Read]:`%s`\n", buf)

	}

	quit <- true

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

	fmt.Println("Wait read to end")
	<-quit
}
