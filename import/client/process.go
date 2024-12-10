package client

import (
	"net"
	"software/import/socket"
)

type Process interface {
	Request(meta *socket.Metadata, server net.Conn) error
	Response(frame *socket.Frame)
}

func (m *Model) UpsertProcess(event string, process Process) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.Processes[event] = process
}
