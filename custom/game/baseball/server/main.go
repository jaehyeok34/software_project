package main

import (
	"log"
	"software/custom/game/baseball/server/server"
	"software/custom/game/baseball/server/server/system/start"
)

func main() {
	model := server.NewModel()
	if err := model.Listen("tcp", "localhost:9999", 2); err != nil {
		log.Fatal(err)
	}
	model.UpsertSystem(start.Event, start.New(model.BaseBall))
	model.Accept()
}
