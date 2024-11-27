package server

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

type PubSubServer struct {
	mu   sync.RWMutex
	subs []net.Conn
}

func NewPubSubServer() *PubSubServer {
	return &PubSubServer{
		subs: make([]net.Conn, 0),
	}
}

func (ps *PubSubServer) Subscribe(conn net.Conn) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ps.subs = append(ps.subs, conn)
}

func (ps *PubSubServer) Update() {
	// 1초마다 모든 구독자에게 메시지를 보냄
	for {
		ps.mu.RLock()
		fmt.Println("메시지를 보냄", len(ps.subs))
		for _, conn := range ps.subs {
			_, err := conn.Write([]byte("hello world" + "\n"))
			if err != nil {
				log.Fatal(err)
			}
		}
		ps.mu.RUnlock()

		time.Sleep(time.Second)
	}
}

func main() {
	ps := NewPubSubServer()

	server, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	fmt.Println("서버 시작")

	go ps.Update()
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("클라이언트 연결됨")
		go ps.Subscribe(conn)
	}
}
