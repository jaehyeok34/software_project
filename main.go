package main

import (
	"fmt"
	"log"
	"net"
	"software/room"
	"software/socket"
)

func main() {
	// room 서버 생성
	room := room.New()
	if err := room.ListenAndServe("tcp", "localhost:9999"); err != nil {
		log.Fatal(err)
	}
	defer room.Server.Close()

	// chatSystem := chat.New()
	// room.AddSystem(chatSystem)
	// room.UpdateSystem(chatSystem, &decorator.TimeStamp{System: chatSystem})

	clientTest()
}

func clientTest() {
	for {
		var input string
		fmt.Print("> ")
		_, err := fmt.Scan(&input)
		if err != nil {
			log.Fatal(err)
		}

		conn, err := net.Dial("tcp", "localhost:9999")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("접속 성공")
		if err := socket.Send(conn, []byte("hello world")); err != nil {
			log.Fatal(err)
		}

		// conn.Close()

		// time.Sleep(time.Second * 3)
		// conn.Close()
	}
}
