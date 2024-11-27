package room

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

type System interface {
	Run(args ...interface{})
}

type Room struct {
	mu      sync.RWMutex
	Server  net.Listener
	clients []net.Conn
	systems map[string]System
}

func New() *Room {
	return &Room{
		Server:  nil,
		clients: make([]net.Conn, 0),
		systems: make(map[string]System),
	}
}

func (r *Room) Append(conn net.Conn) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.clients = append(r.clients, conn)
}

func (r *Room) Listen(network string, address string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	server, err := net.Listen(network, address)
	if err != nil {
		return err
	}

	r.Server = server
	return nil
}

func (r *Room) Serve() {
	go func() {
		for {
			conn, err := r.Server.Accept()
			if err != nil {
				log.Fatal(err)
			}
			r.Append(conn)

			r.mu.RLock()
			fmt.Println("클라이언트 연결됨: ", len(r.clients))
			r.mu.RUnlock()
		}
	}()
}

func (r *Room) Test() {
	go func() {
		for {
			r.mu.RLock()
			for _, client := range r.clients {
				_, err := client.Write([]byte("hello world"))
				if err != nil {
					log.Fatal(err)
				}
			}
			r.mu.RUnlock()

			time.Sleep(time.Second)
		}
	}()
}
