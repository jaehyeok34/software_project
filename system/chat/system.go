package chat

import (
	"fmt"
	"software/server"
	"software/socket"
)

type System struct{}

func New() *System {
	return new(System)
}

func (s *System) Event() string {
	return "chat"
}

func (s *System) Run(src *server.Session, dsts []*server.Session, frame *socket.Frame) {
	var messages []string
	for _, arg := range frame.Args {
		if msg, ok := arg.(string); ok {
			messages = append(messages, msg)
		}
	}

	fmt.Println(src.Meta.Name, "받은 메시지:", messages)
	for _, dst := range dsts {
		socket.Write(dst.Conn, frame)
	}
}

var _ server.System = &System{}
