package room

import (
	"fmt"
	"log"
	"net"
)

func (r *Room) Append(conn net.Conn) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.clients = append(r.clients, conn)
}

func (r *Room) ListenAndServe(network string, address string) error {
	if err := r.Listen(network, address); err != nil {
		return err
	}

	go r.Serve()
	return nil
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
	for {
		conn, err := r.Server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		r.Append(conn)

		r.mu.RLock()
		fmt.Println("클라이언트 추가 함( 현재 클라이언트 수:", len(r.clients), ")")
		r.mu.RUnlock()
	}
}
