package room

import (
	"encoding/json"
	"fmt"
	"net"
	"software/import/client"
	"software/import/socket"
	"sync"
)

type Model struct {
	mu       sync.RWMutex
	Listener net.Listener
	clients  map[string]*Connection
	systems  map[string]System
}

type Connection struct {
	Name string
	Conn net.Conn
}

func New() *Model {
	return &Model{
		Listener: nil,
		clients:  make(map[string]*Connection),
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

		n := new(client.Name)
		decoder := json.NewDecoder(conn)
		if err := decoder.Decode(n); err != nil {
			fmt.Println("연결 실패")
			continue
		}

		m.Append(n.Value, conn)
	}
}

func (m *Model) Append(name string, conn net.Conn) {
	m.mu.Lock()
	defer m.mu.Unlock()

	c := &Connection{Name: name, Conn: conn}
	m.clients[c.Name] = c
	fmt.Println("현재 클라이언트 수:", len(m.clients))

	go m.read(c)
}

func (m *Model) read(conn *Connection) {
	for {
		fmt.Println(conn.Name, "read 하는 중")
		req, err := socket.Read(conn.Conn) // blocking
		if err != nil {
			fmt.Println(conn.Name, "read error:", err)
			delete(m.clients, conn.Name)
			fmt.Println("삭제 후:", len(m.clients))
			return
		}

		m.run(conn, req.Event, req.Args...)
	}
}

func (m *Model) run(src *Connection, key string, args ...interface{}) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if !m.has(key) {
		fmt.Println("process failed: key not found")
		return
	}

	var conns []net.Conn
	for _, conn := range m.clients {
		conns = append(conns, conn.Conn)
	}

	m.systems[key].Run(src, conns, args...)
}
