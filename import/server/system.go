package server

import "software/import/socket"

type System interface {
	Run(src *Session, dsts []*Session, frame *socket.Frame)
}

func (m *Model) UpsertSystem(event string, system System) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.systems[event] = system
}
