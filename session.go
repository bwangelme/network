package main

import "net"

type CoreSessionIdentify struct {
	id int64
}

func (self *CoreSessionIdentify) ID() int64 {
	return self.id
}

func (self *CoreSessionIdentify) SetID(id int64) {
	self.id = id
}

type tcpSession struct {
	CoreSessionIdentify

	// 原始TCP连接
	conn net.Conn
}

func (self *tcpSession) Raw() net.Conn {
	return self.conn
}

type Session interface {
	ID() int64
	Raw() net.Conn
}

func NewSession(conn net.Conn) Session {
	return &tcpSession{
		conn: conn,
	}
}
