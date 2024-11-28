package room

import (
	"net"
	"reflect"
)

type System interface {
	Run(conns []net.Conn, args ...interface{})
}

func (r *Room) AddSystem(system System) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.systems[reflect.TypeOf(system)] = system
	// fmt.Println("system length:", len(r.systems))
}

func (r *Room) Process(systemType System, args ...interface{}) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.systems[reflect.TypeOf(systemType)].Run(r.clients, args...)
}
