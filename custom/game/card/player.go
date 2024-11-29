package card

import (
	"software/import/client"
	"software/import/system/chat"
)

type Player struct {
	*client.Model
}

func New(name string) *Player {
	return &Player{client.New(name)}
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
