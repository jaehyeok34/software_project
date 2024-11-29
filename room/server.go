package room

import (
	"net"
	"reflect"
	"sync"
)

type Server struct {
	mu      sync.RWMutex
	Server  net.Listener
	clients []net.Conn
	systems map[reflect.Type]System
}

func New() *Server {
	return &Server{
		Server:  nil,
		clients: make([]net.Conn, 0),
		systems: make(map[reflect.Type]System),
	}
}
