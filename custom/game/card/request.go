package card

import (
	"fmt"
	"software/import/socket"
	"software/import/system/chat"
)

func (p *Player) SendChat(message string) error {
	p.Mu.RLock()
	defer p.Mu.RUnlock()

	req := new(socket.Frame)
	req.Event = chat.Key
	req.Args = append(req.Args, message)

	err := socket.Write(p.Conn, req)
	if err != nil {
		fmt.Println("Chat Write 문제 발생:", err)
		return err
	}
	return nil
}
