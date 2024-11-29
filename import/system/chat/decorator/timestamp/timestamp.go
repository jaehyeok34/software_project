package timestamp

import (
	"fmt"
	"net"
	"software/import/room"
	"time"
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

func (s *System) Run(conns []net.Conn, args ...interface{}) {
	for i, arg := range args {
		if v, ok := arg.(string); ok {
			args[i] = fmt.Sprintf("[%s] %s", time.Now().Format("15:04:05"), v)
		}
	}
	s.System.Run(conns, args...)
}
