package room

import (
	"fmt"
	"net"
	"software/socket"
	"sync"
)

type Server struct {
	mu      sync.RWMutex
	Server  net.Listener
	clients []net.Conn
	systems map[string]System
}

func New() *Server {
	return &Server{
		Server:  nil,
		clients: make([]net.Conn, 0),
		systems: make(map[string]System),
	}
}

func (s *Server) read(conn net.Conn) {
	req, err := socket.Read(conn)
	if err != nil {
		fmt.Println("read error:", err)
		return
	}

	// TODO: req.Event 확인 후 System 동작
	s.Process(req.Event, req.Args...)
}
