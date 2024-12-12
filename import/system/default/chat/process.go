package chat

import (
	"fmt"
	"software/import/socket"
	"software/import/system"
)

var Event = "chat"

type Process struct{}

func NewProcess() *Process {
	return new(Process)
}

// implementation
// 서버가 "chat" 이벤트를 수신했을 때, 처리하는 로직이다.
func (p *Process) Run(src *socket.Metadata, frame *socket.Frame, sessions []*socket.Session) {
	// 프레임에서 메시지(string)만 추출한다.
	var messages []string
	for _, arg := range frame.Args {
		if msg, ok := arg.(string); ok {
			messages = append(messages, msg)
		}
	}

	// 메시지를 모든 세션에게 재전송한다.
	fmt.Println(src.Name, "받은 메시지:", messages)
	for _, dst := range sessions {
		socket.Write(dst.Conn, frame)
	}
}

var _ system.Process = new(Process)
