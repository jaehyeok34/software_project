package main

import (
	"fmt"
	"log"
	"software/game/card"
	"time"
)

func main() {
	player := card.New()
	if err := player.ConnectAndListen("tcp", "localhost:9999"); err != nil {
		log.Fatal("문제 발생:", err)
	}

	for {
		// ---------------테스트용 코드임 삭제하셈--------------------
		player.Chat("hello world")
		fmt.Println("메시지 보냄")
		time.Sleep(time.Second)
	}
}
