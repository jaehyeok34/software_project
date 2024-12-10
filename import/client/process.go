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
	m.processes.Write(func(storage map[string]Process) {
		storage[event] = process
	})
}

func (m *Model) GetProcess(event string) Process {
	if process, ok := m.processes.Read(func(storage map[string]Process) any {
		return storage[event]
	}).(Process); ok {
		return process
	}

	return nil
}
