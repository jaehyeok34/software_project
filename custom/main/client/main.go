package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"software/custom/game/card"
)

func main() {
	p := card.New("cilent1")
	if err := p.ConnectAndListen("tcp", "localhost:9999"); err != nil {
		log.Fatal("문제 발생:", err)
	}

	go p.Process()

	for {
		// ---------------테스트용 코드임 삭제하셈--------------------
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			p.SendChat(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("스캔에 문제 생김")
		}
	}
}
