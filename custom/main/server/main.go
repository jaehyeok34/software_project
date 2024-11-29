package main

import (
	"log"
	"software/custom/game/card/system/shuffle"
	"software/import/room"
	"software/import/system/chat"
	"software/import/system/chat/decorator/client"
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
	room.AddSystem(shuffle.Key, shuffle.New())

	time.Sleep(time.Second * 5)
	timestamp := timestamp.New(chatSystem)
	room.UpdateSystem(chat.Key, timestamp)

	time.Sleep(time.Second * 3)
	client := client.New(timestamp)
	room.UpdateSystem(chat.Key, client)

	for {
	}
}
