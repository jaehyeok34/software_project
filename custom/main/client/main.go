package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"software/custom/game/card"
	"time"
)

func main() {
	p := card.New("cilent1")
	if err := p.ConnectAndListen("tcp", "localhost:9999"); err != nil {
		log.Fatal("문제 발생:", err)
	}

	go p.Process()

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
			sendChat(p)

		case "2":
			p.Suffle()
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

func sendChat(p *card.Player) {
	fmt.Print("> ")
	message, err := scanner()
	if err != nil {
		fmt.Println("scanner error")
		return
	}

	p.SendChat(message)
	time.Sleep(time.Millisecond * 10)
}
