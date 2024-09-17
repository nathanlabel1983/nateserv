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

	connectionsCh     chan *HTTPConnection // Channel for incoming connections
	activeConnections []*HTTPConnection    // Slice of all active connections
}

func NewServer() *Server {
	s := &Server{
		port:           DefaultPort,
		requestsCh:     make(chan Request, 100),
		requestHandler: defaultRequestHandler,

		connectionsCh:     make(chan *HTTPConnection, 10),
		activeConnections: make([]*HTTPConnection, 0),
	}
	return s
}

func (s *Server) handleConnection(connection net.Conn) {
	fmt.Printf("Connection Received: %s\n", connection.LocalAddr().String())
	defer connection.Close()
	reader := bufio.NewReader(connection)
	writer := bufio.NewWriter(connection)
	d := make([]string, 0)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil || msg == "\r\n" {
			r := newRequest(d)
			fmt.Printf("Sinking Request: %s\n", string(r.Method))
			s.requestsCh <- r
			break
		}
		d = append(d, msg)
	}
	res := NewResponse(StatusOK, "<p>Hello World</p>")
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
			fmt.Println("unable to start TCP Server")
			continue
		}
		go s.handleConnection(c)
	}
}

func defaultRequestHandler(r Request) error {
	if r.Method == "GET" {
		fmt.Println("Processing GET")
	}
	return nil
}
