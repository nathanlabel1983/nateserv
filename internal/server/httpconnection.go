package server

import (
	"bufio"
	"net"
)

type HTTPConnection struct {
	connection net.Conn
	readWriter bufio.ReadWriter
}

func NewHTTPConnection(conn net.Conn) *HTTPConnection {
	h := HTTPConnection{
		connection: conn,
		readWriter: *bufio.NewReadWriter(
			bufio.NewReader(conn),
			bufio.NewWriter(conn),
		),
	}
	return &h
}
