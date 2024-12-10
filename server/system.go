package server

import "software/socket"

type System interface {
	Run(src *Session, dsts []*Session, frame *socket.Frame)
	Event() string
}

func (m *Model) UpsertSystem(system System) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.systems[system.Event()] = system
}
