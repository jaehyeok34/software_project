package main

import (
	"fmt"
	"log"
	"software/custom/game/card"
	"time"
)

func main() {
	player := card.New("cilent1")
	if err := player.ConnectAndListen("tcp", "localhost:9999"); err != nil {
		log.Fatal("문제 발생:", err)
	}

	for {
		// ---------------테스트용 코드임 삭제하셈--------------------
		player.Chat("hello world")
		fmt.Println(player.Name, "메시지 보냄")
		time.Sleep(time.Second)
	}
}
