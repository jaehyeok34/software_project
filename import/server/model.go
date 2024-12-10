package server

import (
	"fmt"
	"io"
	"net"
	"software/import/socket"
	"sync"
)

type Model struct {
	listener net.Listener
	capacity uint16

	mu       sync.RWMutex
	sessions map[string]*Session
	systems  map[string]System
}

func New() *Model {
	return &Model{
		sessions: make(map[string]*Session),
		systems:  make(map[string]System),
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

		if len(m.sessions) >= int(m.capacity) {
			fmt.Println("가득 참")
			conn.Close()
			continue
		}

		meta, err := socket.RecieveConnection(conn)
		if err != nil {
			fmt.Println("연결 실패")
			continue
		}

		m.AddSession(meta, conn)
	}
}

func (m *Model) AddSession(meta *socket.Metadata, conn net.Conn) {
	m.mu.Lock()
	defer m.mu.Unlock()

	session := &Session{Meta: meta, Conn: conn}
	m.sessions[meta.Name] = session
	fmt.Println("현재 클라이언트 수:", len(m.sessions))

	go m.Read(session)
}

func (m *Model) RemoveSession(session *Session) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.sessions, session.Meta.Name)
	fmt.Println("삭제 후", len(m.sessions))
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

		m.Process(session, frame)
	}
}

func (m *Model) Process(session *Session, frame *socket.Frame) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	system := m.systems[frame.Event]
	if system == nil {
		return
	}

	var sessions []*Session
	for _, s := range m.sessions {
		sessions = append(sessions, s)
	}

	system.Run(session, sessions, frame)
}
