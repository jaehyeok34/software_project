package server

import "software/import/socket"

type System interface {
	Run(src *Session, dsts []*Session, frame *socket.Frame)
}

func (m *Model) UpsertSystem(event string, system System) {
	m.systems.Write(func(storage map[string]System) {
		storage[event] = system
	})
}
