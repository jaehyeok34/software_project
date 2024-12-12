package chat

import (
	"fmt"
	"software/import/socket"
	"software/import/system"
	"time"
)

type timestampDecorator struct {
	system.Process
}

func NewTimestampDecorator(chat system.Process) system.Process {
	return &timestampDecorator{chat}
}

func (t *timestampDecorator) Run(src *socket.Metadata, frame *socket.Frame, sessions []*socket.Session) {
	for i, arg := range frame.Args {
		if msg, ok := arg.(string); ok {
			frame.Args[i] = fmt.Sprintf("[%s] %s", time.Now().Format("15:04:05"), msg)
		}
	}

	t.Process.Run(src, frame, sessions)
}

var _ system.Process = new(timestampDecorator)
