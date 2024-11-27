package main

import (
	"fmt"
	"log"
	"net"
)

func connect() (net.Conn, error) {
	conn, err := net.Dial("tcp", "localhost:9999")
	if err != nil {
		return nil, err
	}

	return conn, err
}

func main() {
	conn, err := connect()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Println("연결됨")

	buf := make([]byte, 256)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
			break
		}
		fmt.Println("msg: ", string(buf[:n]))
	}
}
