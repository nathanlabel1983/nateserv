package main

import "github.com/nathanlabel1983/nateserv/internal/server"

func main() {
	s := server.NewServer()
	s.Start()
}
