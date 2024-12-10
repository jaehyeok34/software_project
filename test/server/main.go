package main

import (
	"software/import/default/system/chat"
	"software/import/server"
)

func main() {
	server := server.New()
	server.Listen("tcp", "localhost:9999", 1)
	server.UpsertSystem(chat.Event, chat.New())
	server.Accept()
}
