package base

import (
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
	s.System.Run(src, conns, args...) // chat.System이 호출됨
}

var _ room.System = new(System)
