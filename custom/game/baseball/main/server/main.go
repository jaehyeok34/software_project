package main

import (
	"log"
	"software/custom/game/baseball/model"
	"software/custom/game/baseball/system/notify"
)

func main() {
	model := model.New()
	model.Init()
	if err := model.Listen("tcp", "localhost:9999", 2); err != nil {
		log.Fatal(err)
	}

	// 기능(System) 1. 클라이언트로 부터 notify 라는 이벤트를 받았을 때, args에 있는 숫자를 goal과 비교해서
	// 볼, 스트라이크를 전달해줌
	model.AddSystem("notify", &notify.System{})

	for {

	}
}
