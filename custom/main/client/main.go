package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"software/custom/game/card/player"
	"software/import/client"
	"time"
)

func main() {
	m := client.New(fmt.Sprintf("%s%d", "client", rand.Int()))
	fmt.Println(m.Name)
	if err := m.ConnectAndListen("tcp", "localhost:9999"); err != nil {
		log.Fatal("문제 발생:", err)
	}

	go player.Process(m)

	// UI(CLI 기반)
	for {
		fmt.Println("1. 채팅, 2. 셔플")
		message, err := scanner()
		if err != nil {
			fmt.Println("scanner error")
			return
		}

		switch message {
		case "1":
			sendChat(m)

		case "2":
			player.Shuffle(m)
		}
	}
}

func scanner() (string, error) {
	var message string

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		message = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return message, nil
}

func sendChat(m *client.Model) {
	fmt.Print("> ")
	message, err := scanner()
	if err != nil {
		fmt.Println("scanner error")
		return
	}

	player.SendChat(m, message)
	time.Sleep(time.Millisecond * 10)
}
