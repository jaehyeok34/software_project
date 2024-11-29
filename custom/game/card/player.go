package card

import (
	"fmt"
	"net"
	"software/import/socket"
	"software/import/system/chat"
	"sync"
)

type Player struct {
	Name string
	mu   sync.RWMutex
	Conn net.Conn
}

func New(name string) *Player {
	return &Player{Name: name}
}

func (c *Player) ConnectAndListen(network string, address string) error {
	err := c.Connect(network, address)
	if err != nil {
		return err
	}
	go c.Listen()

	return nil
}

func (c *Player) Connect(network string, address string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	conn, err := net.Dial(network, address)
	if err != nil {
		fmt.Println("연결 실패")
		return err
	}

	c.Conn = conn
	return nil
}

func (c *Player) Listen() {
	for {
		f, err := socket.Read(c.Conn)
		if err != nil {
			fmt.Println("Listen 문제:", err)
		}

		fmt.Println(f.Args)
	}
}

func (c *Player) Chat(message string) error {
	req := new(socket.Frame)
	req.Event = chat.Key
	req.Args = append(req.Args, message)

	err := socket.Write(c.Conn, req)
	if err != nil {
		fmt.Println("Chat Write 문제 발생:", err)
		return err
	}

	return nil
}
