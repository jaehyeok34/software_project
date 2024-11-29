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

func (m *Model) Listen(network string, address string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	server, err := net.Listen(network, address)
	if err != nil {
		return err
	}

	m.Listener = server

	go m.Accept()
	return nil
}

func (m *Model) Accept() {
	for {
		conn, err := m.Listener.Accept()
		if err != nil {
			fmt.Println("클라이언트를 받아들이는 데 문제가 생김", err)
		}

		m.Append(conn)
	}
}

func (m *Model) Append(conn net.Conn) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.clients = append(m.clients, conn)
	fmt.Println("현재 클라이언트 수:", len(m.clients))

	go m.read(conn)
}

func (m *Model) read(conn net.Conn) {
	for {
		req, err := socket.Read(conn) // blocking
		if err != nil {
			fmt.Println("read error:", err)
		}

		m.run(req.Event, req.Args...)
	}
}

func (m *Model) run(key string, args ...interface{}) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if !m.has(key) {
		fmt.Println("process failed: key not found")
		return
	}

	m.systems[key].Run(m.clients, args...)
}
