package answer

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

func NewRequest() system.Request {
	return new(Request)
}

// 서버에 답안 제출 이벤트를 전송한다.
func (r *Request) Send(src *socket.Metadata, dst net.Conn) {
	socket.Write(dst, &socket.Frame{
		Meta:  src,
		Event: Event,
		Args:  append(make([]any, 0), r.getAnswer()),
	})
}

// 서버에 제출한 답안에 대한 결과를 처리한다.
func (r *Request) Process(frame *socket.Frame) {
	for _, arg := range frame.Args {
		if answer, ok := arg.(string); ok {
			fmt.Println(answer)
		}
	}
}

func (r *Request) getAnswer() string {
	var answer string

	scn := bufio.NewScanner(os.Stdin)
	if scn.Scan() {
		answer = scn.Text()
	}

	if err := scn.Err(); err != nil {
		log.Println(err)
		answer = ""
	}

	return answer
}

var _ system.Request = new(Request)
