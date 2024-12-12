package server

import (
	"fmt"
	"net"
	"software/import/collection"
	"software/import/socket"
	"software/import/system"
)

type Model struct {
	listener net.Listener
	capacity uint16

	sessions  *collection.Map[string, *socket.Session]
	processes *collection.Map[string, system.Process]
}

func New() *Model {
	return &Model{
		sessions:  collection.New[string, *socket.Session](),
		processes: collection.New[string, system.Process](),
	}
}

// 서버를 연다.
func (m *Model) Listen(network, address string, capacity uint16) error {
	listener, err := net.Listen(network, address)
	if err != nil {
		return err
	}

	m.listener = listener
	m.capacity = capacity
	return nil
}

// 클라이언트의 연결을 수신한다.
func (m *Model) Accept() {
	for {
		conn, err := m.listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// race condition을 해결하기 위한 로직이다.
		// 현재 session map의 저장된 데이터 개수를 반환받고(int), 허용량보다 크다면 접속을 해제한다.
		// if cap, ok := m.sessions.Read(func(store map[string]*socket.Session) any {
		// 	return len(store) // int
		// }).(int); ok && cap >= int(m.capacity) {
		// 	fmt.Println("가득 참")
		// 	conn.Close()
		// 	continue
		// }
		if m.sessions.Length() >= int(m.capacity) {
			fmt.Println("가득 참")
			conn.Close()
			continue
		}

		// 접속이 허용된 클라이언트로부터 메타데이터를 전달받는다.
		// 메타데이터를 받는 데 오류가 발생하면 접속을 해제한다.
		meta, err := socket.ReceiveMetadata(conn)
		if err != nil {
			fmt.Println("연결 실패", err)
			conn.Close()
			continue
		}

		// 메타데이터의 수신이 완료되면, 세션을 추가한다.
		m.AddSession(meta, conn)
	}
}

func (m *Model) UpsertProcess(event string, process system.Process) {
	m.processes.Store(event, process)
}
