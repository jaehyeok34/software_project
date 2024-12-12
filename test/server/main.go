package main

import (
	"software/import/server"
	"software/import/system/default/chat"
)

func main() {
	server := server.New()
	server.Listen("tcp", "localhost:9999", 2)
	server.UpsertProcess(chat.Event, chat.NewProcess())
	server.Accept()
}
