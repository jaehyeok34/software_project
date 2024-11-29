package main

import (
	"fmt"
	"log"
	"software/custom/game/card"
	"time"
)

func main() {
	p := card.New("cilent1")
	if err := p.ConnectAndListen("tcp", "localhost:9999"); err != nil {
		log.Fatal("문제 발생:", err)
	}

	for {
		// ---------------테스트용 코드임 삭제하셈--------------------
		p.Chat("hello world")
		fmt.Println(p.Name, "메시지 보냄")
		time.Sleep(time.Second)
	}
}
