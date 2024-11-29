package decorator

import (
	"fmt"
	"net"
	"software/room"
	"time"
)

type TimeStamp struct {
	System room.System
}

func (ts *TimeStamp) Run(conns []net.Conn, args ...interface{}) {
	for i, arg := range args {
		if v, ok := arg.(string); ok {
			args[i] = fmt.Sprintf("[%s] %s", time.Now().Format("15:04:05"), v)
		}
	}
	ts.System.Run(conns, args...)
}
