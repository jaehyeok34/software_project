package source

import (
	"fmt"
	"software/import/socket"
	"software/import/system"
)

type Decorator struct {
	system.Process
}

func New(chat system.Process) system.Process {
	return &Decorator{chat}
}

func (d *Decorator) Run(src *socket.Metadata, frame *socket.Frame, sessions []*socket.Session) {
	for i, arg := range frame.Args {
		if msg, ok := arg.(string); ok {
			frame.Args[i] = fmt.Sprintf("%s: %s", src.Name, msg)
		}
	}

	d.Process.Run(src, frame, sessions)
}

var _ system.Process = new(Decorator)
