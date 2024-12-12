package system

import (
	"software/import/socket"
)

// 서버가 클라이언트로부터 받은 이벤트를 처리하는 로직이다.
type Process interface {
	// src: 어떤 클라이언트로부터 받았는지에 대한 정보이다.
	// frame: 받은 데이터이다.
	// sessions: 현재 서버에 연결돼 있는 세션 정보들(처리 로직에서 사용됨) 이다.
	Run(src *socket.Metadata, frame *socket.Frame, sessions []*socket.Session)
}
