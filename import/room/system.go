package room

import (
	"fmt"
	"net"
)

type System interface {
	Run(src net.Conn, conns []net.Conn, args ...interface{})
}

func (m *Model) AddSystem(key string, system System) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.systems[key] = system
}

func (m *Model) UpdateSystem(key string, newSystem System) {
	if !m.has(key) {
		fmt.Println("update failed: key not found")
		return
	}

	m.mu.Lock()
	m.systems[key] = newSystem
	m.mu.Unlock()
}

func (m *Model) has(key string) bool {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if v := m.systems[key]; v != nil {
		return true
	}

	return false
}
