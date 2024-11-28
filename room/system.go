package room

import (
	"fmt"
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
	fmt.Println("system length:", len(r.systems))
}

func (r *Room) UpdateSystem(systemType System, newSystem System) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	key := reflect.TypeOf(systemType)

	is := r.systems[key]
	if is == nil {
		fmt.Println("update failed: not found")
		return
	}

	r.systems[key] = newSystem
}

func (r *Room) Process(systemType System, args ...interface{}) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.systems[reflect.TypeOf(systemType)].Run(r.clients, args...)
}
