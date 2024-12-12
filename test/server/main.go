package main

import (
	"software/import/default/chat"
	"software/import/server"
)

func main() {
	server := server.New()
	server.Listen("tcp", "localhost:9999", 2)
	server.UpsertProcess(chat.Event, chat.NewProcess())
	server.Accept()
}
