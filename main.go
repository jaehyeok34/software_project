package main

import (
	"fmt"
	"log"
	"net"
	"software/room"
	"software/system/base/chat"
	"software/system/base/chat/decorator"
	"time"
)

func main() {
	// room 서버 생성
	room := room.New()
	if err := room.ListenAndServe("tcp", "localhost:9999"); err != nil {
		log.Fatal(err)
	}
	defer room.Server.Close()

	chatSystem := chat.NewChatSystem()
	room.AddSystem(chatSystem)
	room.UpdateSystem(chatSystem, &decorator.TimeStamp{System: chatSystem})

	// client 생성 코드(임시)
	go func() {
		conn, err := net.Dial("tcp", "localhost:9999")
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		buf := make([]byte, 256)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(string(buf[:n]))
		}
	}()

	for {
		room.Process(&chat.System{}, "hello world")
		time.Sleep(time.Second)
	}
}
