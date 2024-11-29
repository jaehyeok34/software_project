package player

import (
	"software/custom/game/card/system/shuffle"
	"software/import/client"
	"software/import/system/chat"
)

func Process(m *client.Model) {
	for {
		f := <-m.Ch
		switch f.Event {
		case chat.Key:
			receiveChat(f.Args)
		case shuffle.Key:
			receiveShuffle(f.Args)
		}
	}
}
