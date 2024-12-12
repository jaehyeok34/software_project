package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
	"software/custom/game/baseball/system/answer"
	"software/custom/game/baseball/system/start"
	"software/import/client"
	"software/import/socket"
	"software/import/system/default/chat"
)

func main() {
	meta := &socket.Metadata{Name: fmt.Sprintf("%s%d", "클라", rand.Int())}
	client := client.New(meta)

	if err := client.ConnectAndListen("tcp", "localhost:9999"); err != nil {
		log.Fatal(err)
	}

	client.UpsertRequest(start.Event, start.NewRequest())
	client.UpsertRequest(answer.Event, answer.NewRequest())
	client.UpsertRequest(chat.Event, chat.NewRequest())

	for {
		if request := client.GetRequest(mainMenu()); request != nil {
			request.Send(client.Meta, client.GetServer())
		}
	}
}

func mainMenu() (event string) {
	fmt.Println("1. 게임 시작, 2. 답안 제출, 3. 채팅")
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return
	}

	switch input {
	case "1":
		return start.Event
	case "2":
		return answer.Event
	case "3":
		return chat.Event
	}

	return ""
}
