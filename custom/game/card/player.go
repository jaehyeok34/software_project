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

// 요청
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

func (p *Player) Process() {
	for {
		f := <-p.Frame
		switch f.Event {
		case chat.Key:
			p.receiveChat(f.Args)
		}
	}
}

func (p *Player) receiveChat(message []interface{}) {
	for _, msg := range message {
		if str, ok := msg.(string); ok {
			fmt.Println("받은 메시지:", str)
		}
	}
}
