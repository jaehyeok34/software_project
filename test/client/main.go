package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"software/client"
	"software/socket"
	"software/system/chat"
)

func main() {
	client := client.New(&socket.Metadata{Name: fmt.Sprintf("%s%d", "클라", rand.Int())})
	client.Connect("tcp", "localhost:9999")
	go client.Listen()

	for {
		var message string
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			message = scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			return
		}

		socket.Write(client.Server, &socket.Frame{
			Meta:  *client.Meta,
			Event: chat.New().Event(),
			Args:  []interface{}{message},
		})
	}
}
