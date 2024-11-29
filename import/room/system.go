package room

import (
	"fmt"
	"net"
)

type System interface {
	Run(conns []net.Conn, args ...interface{})
}

func (r *Model) AddSystem(key string, system System) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.systems[key] = system
	fmt.Println("system length:", len(r.systems))
}

func (r *Model) UpdateSystem(key string, newSystem System) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if !r.has(key) {
		fmt.Println("update failed: key not found")
		return
	}

	r.systems[key] = newSystem
}

func (r *Model) has(key string) bool {
	if v := r.systems[key]; v != nil {
		return true
	}

	return false
}
