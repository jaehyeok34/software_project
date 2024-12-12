package system

import (
	"net"
	"software/import/socket"
)

// 클라이언트가 서버와의 통신(주고/받기)를 요청하는 로직이다.
type Request interface {
	// 서버로 프레임(데이터)을 보내는 로직이다.
	// src: 누가 보내는 지에 대한 정보이다.
	// dst: 어디에 보내는 지(서버 소켓)에 대한 정보이다.
	Send(src *socket.Metadata, dst net.Conn)

	// 서버로부터 전달받은 데이터를 처리하는 로직이다.
	// frame: 서버로부터 받은 데이터이다.
	Process(frame *socket.Frame)
}
