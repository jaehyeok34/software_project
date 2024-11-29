package client

import (
	"fmt"
	"net"
	"software/import/socket"
	"sync"
)

type Model struct {
	Name string
	mu   sync.RWMutex
	Conn net.Conn
}

func New(name string) *Model {
	return &Model{Name: name}
}

func (m *Model) ConnectAndListen(network string, address string) error {
	err := m.Connect(network, address)
	if err != nil {
		return err
	}
	go m.Listen()

	return nil
}

func (m *Model) Connect(network string, address string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	conn, err := net.Dial(network, address)
	if err != nil {
		fmt.Println("연결 실패")
		return err
	}

	m.Conn = conn
	return nil
}

func (c *Model) Listen() {
	for {
		f, err := socket.Read(c.Conn)
		if err != nil {
			fmt.Println("Listen 문제:", err)
		}

		fmt.Println(f.Args)
	}
}