package room

import (
	"fmt"
	"net"
	"reflect"
)

type System interface {
	Run(conns []net.Conn, args ...interface{})
}

func (r *Server) AddSystem(system System) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.systems[reflect.TypeOf(system)] = system
	fmt.Println("system length:", len(r.systems))
}

func (r *Server) UpdateSystem(systemType System, newSystem System) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	key := reflect.TypeOf(systemType)
	if !r.has(key) {
		fmt.Println("update failed: key not found")
		return
	}

	r.systems[key] = newSystem
}

func (r *Server) Process(systemType System, args ...interface{}) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	key := reflect.TypeOf(systemType)
	if !r.has(key) {
		fmt.Println("process failed: key not found")
		return
	}

	r.systems[reflect.TypeOf(systemType)].Run(r.clients, args...)
}

func (r *Server) has(key reflect.Type) bool {
	if v := r.systems[key]; v != nil {
		return true
	}

	return false
}
