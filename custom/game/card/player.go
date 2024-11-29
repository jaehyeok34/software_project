package card

import (
	"fmt"
	"software/import/client"
	"software/import/socket"
	"software/import/system/chat"
)

type Player struct {
	*client.Model
}

func New(name string) *Player {
	return &Player{client.New(name)}
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
