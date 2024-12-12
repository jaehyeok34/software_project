package server

import (
	"fmt"
	"io"
	"net"
	"software/import/socket"
)

// 세션을 추가한다.
func (m *Model) AddSession(meta *socket.Metadata, conn net.Conn) {
	session := &socket.Session{Meta: meta, Conn: conn}

	// race condition 해결을 위한 로직
	// 현재 session map에 연결된 세션을 저장한다.
	// m.sessions.Write(func(store map[string]*socket.Session) {
	// 	store[meta.Name] = session
	// 	fmt.Println("현재 클라이언트 수:", len(store))
	// })
	m.sessions.Store(meta.Name, session)

	// 세션 추가 이후, 통신을 시작한다.
	go m.receive(session)
}

// 세션을 삭제한다.
func (m *Model) removeSession(session *socket.Session) {
	// race condition 해결을 위한 로직
	// 세션을 닫고, session map에서 제거한다.
	// m.sessions.Write(func(store map[string]*socket.Session) {
	// 	session.Conn.Close()
	// 	delete(store, session.Meta.Name)
	// 	fmt.Println("삭제 후", len(store))
	// })
	m.sessions.Delete(session.Meta.Name)
	fmt.Println("삭제 후", m.sessions.Length())
}

// 전달받은 세션과의 통신을 진행한다.
func (m *Model) receive(session *socket.Session) {
	for {
		fmt.Println(session.Meta.Name, "read 하는 중")
		frame, err := socket.Read(session.Conn)
		if err != nil {
			// 연결이 종료되면, 세션을 제거한다.
			if err == io.EOF {
				fmt.Println(session.Meta.Name, "연결 해제")
				m.removeSession(session)
				break
			}

			// 연결 종료가 아닌 다른 에러가 발생했을 때, 해당 에러를 출력한다.
			fmt.Println(session.Meta.Name, "데이터 읽기 에러", err)
			continue
		}

		// 세션으로부터 받은 프레임(데이터)을 처리한다.
		m.processFrame(session.Meta, frame)
	}
}

// 세션으로부터 받은 프레임을 처리한다.
func (m *Model) processFrame(src *socket.Metadata, frame *socket.Frame) {
	// // race condition 해결을 위한 로직
	// // process map에서 세션으로부터 받은 프레임을 처리할 수 있는 프로세스를 획득한다.
	// process, ok := m.processes.Read(func(store map[string]system.Process) any {
	// 	return store[frame.Event]
	// }).(system.Process)

	// // 적절한 프로세스를 획득하지 못한다면, 처리를 중단한다.
	// if !ok {
	// 	return
	// }

	// // 현재 서버에 연결되어 있는 모든 세션 정보를 슬라이스로 획득한다.
	// sessions := m.sessions.Read(func(store map[string]*socket.Session) any {
	// 	// session map의 값들만 추출하고(iter.Seq) 이를 슬라이스로 변환한다.
	// 	return slices.Collect(maps.Values(store))
	// }).([]*socket.Session)

	// // 획득한 프로세스에 처리 로직을 실행한다. (goroutine은 실험적 사용이다.)
	// go process.Run(src, frame, sessions)
	if process, ok := m.processes.Load(frame.Event); ok {
		process.Run(src, frame, m.sessions.GetAll())
	}
}
