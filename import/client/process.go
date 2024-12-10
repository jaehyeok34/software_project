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

	m.processes[event] = process
}

func (m *Model) GetProcess(event string) Process {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.processes[event]
}
