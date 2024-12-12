package main

import (
	"log"
	"software/custom/game/baseball"
	"software/custom/game/baseball/system/answer"
	"software/custom/game/baseball/system/start"
	"software/import/system/default/chat"
)

func main() {
	server := baseball.NewServer(5)
	if err := server.Listen("tcp", "localhost:9999", 2); err != nil {
		log.Fatal(err)
	}

	data := server.GetData()
	server.UpsertProcess(chat.Event, chat.NewProcess())
	server.UpsertProcess(start.Event, start.NewProcess(data))
	server.UpsertProcess(answer.Event, answer.NewProcess(data))

	server.Accept()
}
