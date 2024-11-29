package main

import (
	"log"
	"software/room"
	"software/system/base/chat"
)

func main() {
	room := room.New()
	if err := room.ListenAndServe("tcp", "localhost:9999"); err != nil {
		log.Fatal(err)
	}
	defer room.Server.Close()

	chatSystem := chat.New()
	room.AddSystem("chat", chatSystem)
	// room.UpdateSystem("chat", &decorator.TimeStamp{System: chatSystem})

	for {
	}
}
