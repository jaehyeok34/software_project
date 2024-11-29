package main

import (
	"log"
	suffle "software/custom/game/card/system/shuffle"
	"software/import/room"
	"software/import/system/chat"
	"software/import/system/chat/decorator/timestamp"
	"time"
)

func main() {
	room := room.New()
	if err := room.Listen("tcp", "localhost:9999"); err != nil {
		log.Fatal(err)
	}
	defer room.Listener.Close()

	chatSystem := chat.New()
	room.AddSystem(chat.Key, chatSystem)
	room.AddSystem(suffle.Key, suffle.New())

	time.Sleep(time.Second * 5)
	room.UpdateSystem(chat.Key, timestamp.New(chatSystem))

	for {
	}
}
