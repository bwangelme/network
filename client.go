package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

type TcpClient struct {
	addr string
	quit chan bool
}

func NewTcpClient(addr string) *TcpClient {
	return &TcpClient{
		addr: addr,
		quit: make(chan bool),
	}
}

func (self *TcpClient) write(conn net.Conn) {
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

func (self *TcpClient) read(conn net.Conn) {
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

	self.quit <- true

}

func (self *TcpClient) Start() {
	conn, err := net.Dial("tcp", self.addr)
	if err != nil {
		log.Fatalln(err)
	}

	go self.read(conn)
	self.write(conn)

	fmt.Println("Wait read to end")
	<-self.quit
}
