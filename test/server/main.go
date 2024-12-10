package main

import (
	"software/server"
	"software/system/chat"
)

func main() {
	server := server.New()
	server.Listen("tcp", "localhost:9999", 1)
	server.UpsertSystem(chat.New())

	server.Accept()
}
