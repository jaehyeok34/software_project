package client

import (
	"fmt"
	"net"
	"software/import/room"
)

type System struct {
	System room.System
}

func New(system room.System) *System {
	if system == nil {
		panic("system 필요")
	}

	return &System{System: system}
}

func (s *System) Run(src *room.Connection, conns []net.Conn, args ...interface{}) {
	for i, arg := range args {
		if v, ok := arg.(string); ok {
			args[i] = fmt.Sprintf("%s: %s", src.Name, v)
		}
	}
	s.System.Run(src, conns, args...)
}

var _ room.System = new(System)
