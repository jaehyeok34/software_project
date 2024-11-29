package chat

import (
	"fmt"
	"net"
	"software/import/room"
	"software/import/socket"
)

type System struct{}

var Key string = "chat"

func New() *System {
	return &System{}
}

func (s *System) Run(src net.Conn, conns []net.Conn, args ...interface{}) {
	if len(conns) == 0 {
		fmt.Println("ChatSystem(Run): missing []net.Conn")
		return
	}

	var message string
	for _, arg := range args {
		if v, ok := arg.(string); ok {
			message = v
		}
	}

	if message == "" {
		fmt.Println("ChatSystem(Run): missing message(string)")
	}

	fmt.Println("received:", message)
	for _, conn := range conns {
		f := new(socket.Frame)
		f.Event = Key
		f.Args = append(f.Args, message)
		socket.Write(conn, f)
	}
}

var _ room.System = new(System)
