package server

import (
	"bufio"
	"fmt"
	"net"
)

const (
	DefaultPort = ":8080"
)

type RequestHandler func(r Request) error

type Server struct {
	port string // The Port that the server will listen on, defaults to 8080

	requestsCh     chan Request   // All generated Requests
	requestHandler RequestHandler // Function that is used to handle requests
}

func NewServer() *Server {
	s := &Server{
		port:       DefaultPort,
		requestsCh: make(chan Request),
	}
	return s
}

func (s *Server) handleConnection(connection net.Conn) {
	defer connection.Close()
	reader := bufio.NewReader(connection)
	d := make([]string, 0)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			r := newRequest(d)
			s.requestsCh <- r
			return
		}
		d = append(d, msg)
	}
}

func (s *Server) Start() {
	if s.port == "" {
		s.port = ":8080"
	}
	l, err := net.Listen("tcp", s.port)
	if err != nil {
		fmt.Println("unable to start server")
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
