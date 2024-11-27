package main

import (
	"fmt"
	"log"
	"net"
	"software/room"
)

func main() {
	room := room.New()
	if err := room.Listen("tcp", "localhost:9999"); err != nil {
		log.Fatal(err)
	}
	defer room.Server.Close()

	room.Serve()
	room.Test()

	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("클라이언트 하나 연결 시도됨")
	buf := make([]byte, 256)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(buf[:n]))
	}

}
