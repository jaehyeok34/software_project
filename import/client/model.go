package client

import (
	"fmt"
	"io"
	"net"
	"software/import/collection"
	"software/import/socket"
)

type Model struct {
	Meta      *socket.Metadata
	server    net.Conn
	processes *collection.Map[string, Process]
}

func New(meta *socket.Metadata) *Model {
	return &Model{
		Meta:      meta,
		processes: collection.NewMap[string, Process](),
	}
}

func (m *Model) ConnectAndListen(network, address string) error {
	if err := m.Connect(network, address); err != nil {
		return err
	}

	go m.Listen()
	return nil
}

func (m *Model) Connect(network, address string) error {
	server, err := net.Dial(network, address)
	if err != nil {
		return err
	}

	m.server = server
	return socket.RequestConnection(server, m.Meta)
}

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

		if process := m.GetProcess(frame.Event); process != nil {
			process.Response(frame)
		}
	}
}

func (m *Model) GetServer() net.Conn {
	return m.server
}
