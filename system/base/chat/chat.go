package chat

import (
	"fmt"
	"log"
	"net"
	"software/room"
)

type System struct{}

func New() *System {
	return &System{}
}

func (cs *System) Run(conns []net.Conn, args ...interface{}) {
	fmt.Println("---ChatSystem(Run)---")
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

	for i, conn := range conns {
		_, err := conn.Write([]byte(message))
		if err != nil {
			log.Fatal("ChatSystem(Run):", err, "index:", i)
		}

	}
}

var _ room.System = &System{}
