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
		port:           DefaultPort,
		requestsCh:     make(chan Request, 100),
		requestHandler: defaultRequestHandler,
	}
	return s
}

func (s *Server) handleConnection(connection net.Conn) {
	defer connection.Close()
	reader := bufio.NewReader(connection)
	writer := bufio.NewWriter(connection)
	d := make([]string, 0)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			r := newRequest(d)
			s.requestsCh <- r
			break
		}
		d = append(d, msg)
	}
	res := NewResponse(StatusOK, StatusOKReason, "<p>Hello World</p>")
	writer.WriteString(res.GetResponseString())
	writer.Flush()
}

func (s *Server) handleRequests() {
	for {
		req := <-s.requestsCh
		s.requestHandler(req)
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
	go s.handleRequests()

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

func defaultRequestHandler(r Request) error {
	fmt.Printf("Method: %s\n", r.Method)
	fmt.Printf("Path: %s\n", r.Path)
	fmt.Printf("HttpVersion: %s\n", r.HttpVersion)
	fmt.Println("")
	fmt.Println("Headers")
	for k, v := range r.Headers {
		fmt.Printf("[%s]%s", k, v)
	}
	return nil
}
