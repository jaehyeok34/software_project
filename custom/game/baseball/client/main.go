package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"software/custom/game/baseball/client/client"
	"software/custom/game/baseball/client/client/process/start"
	"software/import/default/process/chat"
)

var event map[string]string = make(map[string]string)

func init() {
	event["1"] = chat.Event
	event["2"] = start.Event
}

func main() {
	model := client.New()
	if err := model.ConnectAndListen("tcp", "localhost:9999"); err != nil {
		log.Fatal(err)
	}

	model.UpsertProcess(chat.Event, chat.New())

	for {
		fmt.Println("1, 채팅, 2. 시작")
		var input string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input = scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			return
		}

		process := model.GetProcess(event[input])
		if process == nil {
			fmt.Println("처리 불가능한 기능")
			continue
		}

		process.Request(model.Meta, model.GetServer())
	}
}
