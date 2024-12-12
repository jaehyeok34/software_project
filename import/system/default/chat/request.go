package chat

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"software/import/socket"
	"software/import/system"
)

type Request struct{}

func NewRequest() *Request {
	return new(Request)
}

// implementation
// 클라이언트가 서버로 채팅 메시지를 보내는 로직이다.
func (r *Request) Send(src *socket.Metadata, dst net.Conn) {
	socket.Write(dst, &socket.Frame{
		Meta:  src,
		Event: Event,
		Args:  append(make([]any, 0), r.getMessage()),
	})
}

// implementation
// 서버로부터 채팅 메시지의 처리 결과를 받아 처리하는 로직이다.
func (r *Request) Process(frame *socket.Frame) {
	for _, arg := range frame.Args {
		if msg, ok := arg.(string); ok {
			fmt.Println("받은 메시지:", msg)
		}
	}
}

// 키보드(os.Stdin)로부터 문자열을 입력받아 반환한다.
func (r *Request) getMessage() string {
	var msg string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		msg = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
		msg = ""
	}

	return msg
}

var _ system.Request = new(Request)
