package card

import (
	"fmt"
	"net"
	"sync"
)

type Player struct {
	ID    int
	Cards []string

	mu   sync.RWMutex
	conn net.Conn
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

	c.conn = conn
	return nil
}

func (c *Player) SendChat() {
	go func() {
		for {
			var input string
			fmt.Print("> ")
			fmt.Scan(&input)

			c.mu.RLock()
			_, err := c.conn.Write([]byte(input))
			if err != nil {
				fmt.Println("Player/SendChat():", err)
			}
			c.mu.RUnlock()
		}
	}()
}
