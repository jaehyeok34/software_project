package player

import (
	"fmt"
	"software/custom/game/card/system/shuffle"
	"software/import/client"
	"software/import/socket"
	"software/import/system/chat"
)

func SendChat(m *client.Model, message string) error {
	m.Mu.RLock()
	defer m.Mu.RUnlock()

	f := new(socket.Frame)
	f.Name = m.Name.Value
	f.Event = chat.Key
	f.Args = append(f.Args, message)

	err := socket.Write(m.Conn, f)
	if err != nil {
		fmt.Println("Chat Write 문제 발생:", err)
		return err
	}

	return nil
}

func Shuffle(m *client.Model) error {
	m.Mu.RLock()
	defer m.Mu.RUnlock()

	f := new(socket.Frame)
	f.Name = m.Name.Value
	f.Event = shuffle.Key

	err := socket.Write(m.Conn, f)
	if err != nil {
		fmt.Println("Shuffle Write 문제 발생:", err)
		return err
	}

	return nil
}
