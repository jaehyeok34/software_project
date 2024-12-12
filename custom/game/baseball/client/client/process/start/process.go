package start

import (
	"net"
	"software/import/client"
	"software/import/socket"
)

type Process struct{}

var Event = "start"

func (p *Process) Request(meta *socket.Metadata, server net.Conn) error {
	panic("null")
}

// Response implements client.Process.
func (p *Process) Response(frame *socket.Frame) {
	panic("unimplemented")
}

var _ client.Process = new(Process)
