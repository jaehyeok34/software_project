package start

import (
	"fmt"
	"net"
	"software/import/socket"
	"software/import/system"
)

type Request struct{}

func NewRequest() *Request {
	return new(Request)
}

// 서버에 게임 시작을 요청하는 로직이다.
func (r *Request) Send(src *socket.Metadata, dst net.Conn) {
	socket.Write(dst, &socket.Frame{
		Meta:  src,
		Event: Event,
	})
}

// 게임 시작 이벤트에 대한 서버의 응답을 처리하는 로직이다.
func (r *Request) Process(frame *socket.Frame) {
	for _, arg := range frame.Args {
		if v, ok := arg.(string); ok {
			fmt.Println(v)
		}
	}
}

var _ system.Request = new(Request)
