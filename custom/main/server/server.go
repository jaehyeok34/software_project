package main

import (
	"log"
	"software/import/room"
	"software/import/system/chat"
	"software/import/system/chat/decorator/timestamp"
	"time"
)

func main() {
	room := room.New()
	if err := room.ListenAndServe("tcp", "localhost:9999"); err != nil {
		log.Fatal(err)
	}
	defer room.Listener.Close()

	chatSystem := chat.New()
	room.AddSystem(chat.Key, chatSystem)

	time.Sleep(time.Second * 5)
	room.UpdateSystem(chat.Key, timestamp.New(chatSystem))

	for {
	}
}
