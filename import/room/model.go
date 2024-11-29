package room

import (
	"fmt"
	"net"
	"software/import/socket"
	"sync"
)

type Model struct {
	mu       sync.RWMutex
	Listener net.Listener
	clients  []net.Conn
	systems  map[string]System
}

func New() *Model {
	return &Model{
		Listener: nil,
		clients:  make([]net.Conn, 0),
		systems:  make(map[string]System),
	}
}

func (m *Model) read(conn net.Conn) {
	req, err := socket.Read(conn)
	if err != nil {
		fmt.Println("read error:", err)
		return
	}

	// TODO: req.Event 확인 후 System 동작
	m.Process(req.Event, req.Args...)
}

func (m *Model) Process(key string, args ...interface{}) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if !m.has(key) {
		fmt.Println("process failed: key not found")
		return
	}

	m.systems[key].Run(m.clients, args...)
}
