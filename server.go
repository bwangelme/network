package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
	"sync/atomic"
)

type TcpServer struct {
	addr        string
	sessManager CoreSessionManager
}

var initial uint32
var mu sync.Mutex
var server *TcpServer

/*
 * NewTcpServer
 *
 * 用法: 外部程序必须先调用 NewTcpServer，然后才能给获取当前模块的tcpserver
 */
func NewTcpServer(addr string) *TcpServer {
	if atomic.LoadUint32(&initial) == 0 {
		mu.Lock()
		defer mu.Unlock()

		if initial == 0 {
			server = &TcpServer{
				addr:        addr,
				sessManager: CoreSessionManager{},
			}
			atomic.StoreUint32(&initial, 1)
		}
	}

	return server
}

func GetTcpServer() *TcpServer {
	return server
}

func (self *TcpServer) ListenAndServe(handler func(net.Conn)) {
	l, err := net.Listen("tcp", self.addr)
	if err != nil {
		log.Fatalln(err)
	}

	defer l.Close()
	fmt.Printf("Listening on `%s`\n", self.addr)

	for {
		conn, err := l.Accept()
		fmt.Printf("Accept Conn from %s\n", conn.RemoteAddr())
		if err != nil {
			log.Fatalln(err)
		}

		go self.core_handler(conn, handler)
	}
}

func (self *TcpServer) core_handler(conn net.Conn, handler func(net.Conn)) {
	sess := NewSession(conn)
	self.sessManager.Add(sess)

	handler(conn)

	self.sessManager.Remove(sess)
}

func conn_handler(conn net.Conn) {
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

	fmt.Printf("Total Copy %d bytes from client\n", total)
	conn.Close()
}
