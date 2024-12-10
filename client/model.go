package client

import (
	"fmt"
	"io"
	"net"
	"software/socket"
)

type Model struct {
	Meta    *socket.Metadata
	Server  net.Conn
	FrameCh chan *socket.Frame
}

func New(meta *socket.Metadata) *Model {
	return &Model{
		Meta:    meta,
		FrameCh: make(chan *socket.Frame),
	}
}

func (m *Model) Connect(network, address string) error {
	server, err := net.Dial(network, address)
	if err != nil {
		return err
	}

	m.Server = server
	return socket.RequestConnection(server, m.Meta)
}

func (m *Model) Listen() {
	for {
		frame, err := socket.Read(m.Server)
		if err != nil {
			if err == io.EOF {
				m.Server.Close()
				break
			}

			fmt.Println("서버로 부터 읽기 분제 발생", err)
			continue
		}

		m.FrameCh <- frame
	}
}
