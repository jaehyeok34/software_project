package card

import (
	suffle "software/custom/game/card/system/shuffle"
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
		f := <-p.Ch
		switch f.Event {
		case chat.Key:
			p.receiveChat(f.Args)
		case suffle.Key:
			p.receiveShuffle(f.Args)
		}
	}
}
