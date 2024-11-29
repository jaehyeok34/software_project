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
	clients  map[uint]net.Conn
	systems  map[string]System
}

var index uint = 0

func New() *Model {
	return &Model{
		Listener: nil,
		clients:  make(map[uint]net.Conn),
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

	index++
	m.clients[index] = conn
	fmt.Println("현재 클라이언트 수:", len(m.clients))

	go m.read(index, conn)
}

func (m *Model) read(key uint, conn net.Conn) {
	for {
		fmt.Println(key, "read 하는 중")
		req, err := socket.Read(conn) // blocking
		if err != nil {
			fmt.Println(key, "read error:", err)
			delete(m.clients, key)
			fmt.Println("삭제 후:", len(m.clients))
			return
		}

		m.run(conn, req.Event, req.Args...)
	}
}

func (m *Model) run(src net.Conn, key string, args ...interface{}) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if !m.has(key) {
		fmt.Println("process failed: key not found")
		return
	}

	var conns []net.Conn
	for _, conn := range m.clients {
		conns = append(conns, conn)
	}

	m.systems[key].Run(src, conns, args...)
}
