package card

import (
	"fmt"
	"net"
	"software/socket"
	"sync"
)

type Player struct {
	mu   sync.RWMutex
	Conn net.Conn
}

func NewPlayer() *Player {
	return &Player{}
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

func (c *Player) Chat(message string) error {
	req := new(socket.Frame)
	req.Event = "chat"
	req.Args = append(req.Args, message)

	err := socket.Write(c.Conn, req)
	if err != nil {
		fmt.Println("Chat Write 문제 발생:", err)
		return err
	}

	return nil
}
