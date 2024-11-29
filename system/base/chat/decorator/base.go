package decorator

import (
	"net"
	"software/room"
)

type Base struct {
	System room.System
}

func (b *Base) Run(conns []net.Conn, args ...interface{}) {
	b.System.Run(conns, args...) // chat.System이 호출됨
}

var _ room.System = &Base{}
