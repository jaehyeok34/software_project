package start

import (
	"fmt"
	"software/custom/game/baseball"
	"software/import/socket"
	"software/import/system"
)

var Event = "start"

type process struct {
	data *baseball.Data
}

func NewProcess(data *baseball.Data) system.Process {
	return &process{data}
}

// 클라이언트로부터 시작 이벤트를 전달 받으면 이를 처리하는 로직이다.
func (p *process) Run(src *socket.Metadata, frame *socket.Frame, sessions []*socket.Session) {
	msg := "게임이 시작됐습니다. 숫자를 입력해 주세요"
	if !p.data.Init() {
		msg = "게임이 이미 진행 중 입니다."
	}

	fmt.Println("답:", p.data.GetGoal())
	frame.Args = append(make([]any, 0), msg)
	for _, session := range sessions {
		socket.Write(session.Conn, frame)
	}
}

var _ system.Process = new(process)
