package chat

import (
	"fmt"
	"software/import/socket"
	"software/import/system"
)

type sourceDecorator struct {
	system.Process
}

func NewSourceDecorator(chat system.Process) system.Process {
	return &sourceDecorator{chat}
}

func (s *sourceDecorator) Run(src *socket.Metadata, frame *socket.Frame, sessions []*socket.Session) {
	for i, arg := range frame.Args {
		if msg, ok := arg.(string); ok {
			frame.Args[i] = fmt.Sprintf("%s: %s", src.Name, msg)
		}
	}

	s.Process.Run(src, frame, sessions)
}

var _ system.Process = new(sourceDecorator)
