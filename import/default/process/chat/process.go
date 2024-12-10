package chat

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"software/import/client"
	"software/import/default/system/chat"
	"software/import/socket"
)

type Process struct{}

var Event = "chat"

func New() *Process {
	return new(Process)
}

func (p *Process) Request(meta *socket.Metadata, server net.Conn) error {
	frame := new(socket.Frame)
	frame.Meta = meta
	frame.Event = chat.Event
	frame.Args = make([]interface{}, 0)

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		frame.Args = append(frame.Args, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return socket.Write(server, frame)
}

func (p *Process) Response(frame *socket.Frame) {
	for _, arg := range frame.Args {
		if msg, ok := arg.(string); ok {
			fmt.Println("받은 메시지:", msg)
		}
	}
}

var _ client.Process = new(Process)
