package socket

import "net"

// 소켓의 정보와 실제 연결된 소켓을 담는 구조체이다.(주로 서버에서 사용됨)
type Session struct {
	Meta *Metadata `json:"meta"`
	Conn net.Conn  `json:"conn"`
}
