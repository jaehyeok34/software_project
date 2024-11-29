package room

import (
	"fmt"
	"net"
	"time"
)

func (s *Server) ListenAndServe(network string, address string) error {
	if err := s.Listen(network, address); err != nil {
		return err
	}

	go s.Serve()
	return nil
}

func (s *Server) Listen(network string, address string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	server, err := net.Listen(network, address)
	if err != nil {
		return err
	}

	s.Server = server

	go s.Accept()

	return nil
}

func (s *Server) Accept() {
	for {
		s.mu.RLock()
		conn, err := s.Server.Accept()
		if err != nil {
			fmt.Println("클라이언트를 받아들이는 데 문제가 생김", err)
		}
		s.mu.RUnlock()

		s.Append(conn)
	}
}

func (s *Server) Serve() {
	for {
		s.mu.RLock()
		for i, client := range s.clients {
			go s.read(client)
			fmt.Println(i+1, "번 쨰 client 데이터 읽는 중")
		}
		s.mu.RUnlock()

		time.Sleep(time.Second)
	}
}

func (s *Server) Append(conn net.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.clients = append(s.clients, conn)
	fmt.Println("현재 클라이언트 수:", len(s.clients))
}
