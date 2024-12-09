package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"software/import/client"
	"software/import/socket"
	"time"
)

func main() {
	// 서버에 참가하기
	baseballClient := client.New(fmt.Sprintf("%s%d", "플레이어", rand.Int()))
	baseballClient.ConnectAndListen("tcp", "localhost:9999")

	for {
		run(baseballClient)
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

// 키보드로 부터 값을 받아오는 과정
func run(m *client.Model) {
	fmt.Print("> ")
	message, err := scanner()
	if err != nil {
		fmt.Println("scanner error")
		return
	}

	notify(m, message)
	time.Sleep(time.Millisecond * 10)
}

// 실제 서버로 데이터를 전송하는 과정
func notify(m *client.Model, message string) error {
	m.Mu.RLock()
	defer m.Mu.RUnlock()

	f := new(socket.Frame)
	f.Name = m.Name.Value
	f.Event = "notify"
	f.Args = append(f.Args, message)

	return socket.Write(m.Conn, f)
}
