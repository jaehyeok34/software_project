package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"software/import/client"
	"software/import/default/chat"
	"software/import/socket"
)

func main() {
	client := client.New(&socket.Metadata{Name: fmt.Sprintf("%s%d", "클라", rand.Int())})
	if err := client.ConnectAndListen("tcp", "localhost:9999"); err != nil {
		log.Fatal(err)
	}

	client.UpsertRequest(chat.Event, chat.NewRequest())

	for {
		fmt.Println("1. 채팅")
		var input string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input = scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			return
		}

		var event string
		switch input {
		case "1":
			event = chat.Event
		}

		if request := client.GetRequest(event); request != nil {
			request.Send(client.Meta, client.GetServer())
		}
	}
}
