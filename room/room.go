package room

import (
	"net"
	"reflect"
	"sync"
)

type Room struct {
	mu      sync.RWMutex
	Server  net.Listener
	clients []net.Conn
	// systems map[string]System
	systems map[reflect.Type]System
}

func New() *Room {
	return &Room{
		Server:  nil,
		clients: make([]net.Conn, 0),
		systems: make(map[reflect.Type]System),
	}
}
