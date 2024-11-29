package chat

import (
	"fmt"
	"net"
	"software/room"
	"software/socket"
)

type System struct{}

func New() *System {
	return &System{}
}

func (cs *System) Run(conns []net.Conn, args ...interface{}) {
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
		res := new(socket.Frame)
		res.Args = append(res.Args, message)
		socket.Write(conn, res)
	}
}

var _ room.System = &System{}
