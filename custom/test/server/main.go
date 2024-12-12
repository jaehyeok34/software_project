package main

import (
	"fmt"
	"software/import/server"
	"software/import/system/default/chat"
	"software/import/system/default/chat/decorator/timestamp"
	"time"
)

func main() {
	server := server.New()
	server.Listen("tcp", "localhost:9999", 2)

	chatEvent := chat.NewProcess()
	server.UpsertProcess(chat.Event, chatEvent)

	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("timestamp 기능 추가")

		chatEvent = timestamp.New(chatEvent)
		server.UpsertProcess(chat.Event, chatEvent)
	}()
	server.Accept()
}
