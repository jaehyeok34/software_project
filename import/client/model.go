package client

import (
	"encoding/json"
	"fmt"
	"net"
	"software/import/socket"

	"sync"
)

type Name struct {
	Value string `json:"value"`
}

type Model struct {
	Name *Name
	Mu   sync.RWMutex
	Conn net.Conn
	Ch   chan *socket.Frame
}

func New(name string) *Model {
	return &Model{
		Name: &Name{Value: name},
		Ch:   make(chan *socket.Frame),
	}
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
	m.Mu.Lock()
	defer m.Mu.Unlock()

	conn, err := net.Dial(network, address)
	if err != nil {
		fmt.Println("연결 실패")
		return err
	}

	m.Conn = conn

	e := json.NewEncoder(conn)
	if err := e.Encode(m.Name); err != nil {
		fmt.Println("json encode error:", err)
		return err
	}

	return nil
}

func (m *Model) Listen() {
	for {
		f, err := socket.Read(m.Conn)
		if err != nil {
			fmt.Println("Listen 문제:", err)
		}

		m.Ch <- f
	}
}
