package server

import (
	"fmt"
	"io"
	"net"
	"software/import/collection"
	"software/import/socket"
)

type Model struct {
	listener net.Listener
	capacity uint16

	sessions *collection.Map[string, *Session]
	systems  *collection.Map[string, System]
}

func New() *Model {
	return &Model{
		sessions: collection.NewMap[string, *Session](),
		systems:  collection.NewMap[string, System](),
	}
}

func (m *Model) Listen(network, address string, capacity uint16) error {
	listener, err := net.Listen(network, address)
	if err != nil {
		return err
	}

	m.listener = listener
	m.capacity = capacity
	return nil
}

func (m *Model) Accept() {
	for {
		conn, err := m.listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		if cap, ok := m.sessions.Read(func(storage map[string]*Session) any {
			return len(storage)
		}).(int); ok && cap >= int(m.capacity) {
			fmt.Println("가득 참")
			conn.Close()
			continue
		}

		meta, err := socket.RecieveConnection(conn)
		if err != nil {
			fmt.Println("연결 실패", err)
			continue
		}

		m.AddSession(meta, conn)
	}
}

func (m *Model) AddSession(meta *socket.Metadata, conn net.Conn) {
	session := &Session{Meta: meta, Conn: conn}

	m.sessions.Write(func(storage map[string]*Session) {
		storage[meta.Name] = session
		fmt.Println("현재 클라이언트 수:", len(storage))
	})

	go m.Read(session)
}

func (m *Model) RemoveSession(session *Session) {
	m.sessions.Write(func(storage map[string]*Session) {
		delete(storage, session.Meta.Name)
		fmt.Println("삭제 후", len(storage))
	})
}

func (m *Model) Read(session *Session) {
	for {
		fmt.Println(session.Meta.Name, "read 하는 중")
		frame, err := socket.Read(session.Conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println(session.Meta.Name, "연결 해제")
				m.RemoveSession(session)
				break
			}

			fmt.Println(session.Meta.Name, "데이터 읽기 에러", err)
			continue
		}

		m.Run(session, frame)
	}
}

func (m *Model) Run(session *Session, frame *socket.Frame) {
	system, ok := m.systems.Read(func(storage map[string]System) any {
		return storage[frame.Event]
	}).(System)

	if !ok {
		return
	}

	sessions := m.sessions.Read(func(storage map[string]*Session) any {
		var sessions []*Session
		for _, session := range storage {
			sessions = append(sessions, session)
		}

		return sessions
	}).([]*Session)

	system.Run(session, sessions, frame)
}
