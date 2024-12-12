package client

import (
	"fmt"
	"io"
	"net"
	"software/import/collection"
	"software/import/socket"
	"software/import/system"
)

type Model struct {
	Meta     *socket.Metadata
	server   net.Conn
	requests *collection.Map[string, system.Request]
}

func New(meta *socket.Metadata) *Model {
	return &Model{
		Meta:     meta,
		requests: collection.NewMap[string, system.Request](),
	}
}

// 클라이언트가 서버에 연결하고, 통신을 시작한다.
func (m *Model) ConnectAndListen(network, address string) error {
	if err := m.Connect(network, address); err != nil {
		return err
	}

	go m.Listen()
	return nil
}

// 클라이언트가 서버에 연결을 시도한다.
func (m *Model) Connect(network, address string) error {
	server, err := net.Dial(network, address)
	if err != nil {
		return err
	}

	m.server = server
	return socket.SendMetadata(server, m.Meta)
}

// 서버와 통신을 시작한다. 데이터를 수신하면, 적절한 처리 함수(system.Request.Process)를 호출한다.
func (m *Model) Listen() {
	for {
		frame, err := socket.Read(m.server)
		if err != nil {
			if err == io.EOF {
				m.server.Close()
				break
			}

			fmt.Println("서버로 부터 읽기 분제 발생", err)
			continue
		}

		// 적절한 처리함수를 조회하고, 있다면 호출한다.
		if request := m.GetRequest(frame.Event); request != nil {
			go request.Process(frame)
		}
	}
}

// Request 구현체를 추가한다.
func (m *Model) UpsertRequest(event string, request system.Request) {
	m.requests.Store(event, request)
}

// 외부 패키지에서 이벤트를 통해 Request 객체를 획득하는 로직이다.
func (m *Model) GetRequest(event string) system.Request {
	if request, ok := m.requests.Load(event); ok {
		return request
	}

	return nil
}

func (m *Model) GetServer() net.Conn {
	return m.server
}
