package main

import (
	"fmt"
	"software/import/server"
	"software/import/system/default/chat"
	"time"
)

func main() {
	server := server.New()
	server.Listen("tcp", "localhost:9999", 2)

	chatEvent := chat.NewProcess()
	server.UpsertProcess(chat.Event, chatEvent)

	go func() {
		time.Sleep(time.Second * 5)
		fmt.Println("chat 기능 확장(timestamp)")
		chatEvent = chat.NewTimestampDecorator(chatEvent)
		server.UpsertProcess(chat.Event, chatEvent)

		time.Sleep(time.Second * 5)
		fmt.Println("chat 기능 확장(source)")
		chatEvent = chat.NewSourceDecorator(chatEvent)
		server.UpsertProcess(chat.Event, chatEvent)
	}()

	server.Accept()
}
