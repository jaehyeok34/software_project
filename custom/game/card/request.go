package card

import (
	"fmt"
	suffle "software/custom/game/card/system/shuffle"
	"software/import/socket"
	"software/import/system/chat"
)

func (p *Player) SendChat(message string) error {
	p.Mu.RLock()
	defer p.Mu.RUnlock()

	f := new(socket.Frame)
	f.Event = chat.Key
	f.Args = append(f.Args, message)

	err := socket.Write(p.Conn, f)
	if err != nil {
		fmt.Println("Chat Write 문제 발생:", err)
		return err
	}

	return nil
}

func (p *Player) Suffle() error {
	p.Mu.RLock()
	defer p.Mu.RUnlock()

	f := new(socket.Frame)
	f.Event = suffle.Key

	err := socket.Write(p.Conn, f)
	if err != nil {
		fmt.Println("Suffle Write 문제 발생:", err)
		return err
	}

	return nil
}
