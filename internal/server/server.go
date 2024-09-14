package server

import (
	"bufio"
	"fmt"
	"net"
)

type Server struct {
	port string // The Port that the server will listen on, defaults to 8080
}

func NewServer() *Server {
	s := &Server{
		port: ":8080",
	}
	return s
}

func (s *Server) handleConnection(connection net.Conn) {
	defer connection.Close()
	reader := bufio.NewReader(connection)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("connection closed")
			return
		}
		fmt.Printf("Message Received: %s", msg)
	}
}

func (s *Server) Start() {
	if s.port == "" {
		s.port = ":8080"
	}
	l, err := net.Listen("tcp", s.port)
	if err != nil {
		fmt.Errorf("unable to start server")
		return
	}
	defer l.Close()
	fmt.Printf("TCP server listening on port %s\n", s.port)
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("error accepting connection")
			continue
		}
		go s.handleConnection(c)
	}
}
