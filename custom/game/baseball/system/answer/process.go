package answer

import (
	"fmt"
	"software/custom/game/baseball"
	"software/import/socket"
	"software/import/system"
	"strings"
)

var Event = "answer"

type Process struct {
	data *baseball.Data
}

func NewProcess(data *baseball.Data) *Process {
	return &Process{data}
}

// 클라이언트로부터 답안 제출 이벤트를 수신했을 때 처리하는 로직이다.
func (p *Process) Run(src *socket.Metadata, frame *socket.Frame, sessions []*socket.Session) {
	// 수신한 프레임에서 답안(string)을 추출한다.
	var answer string
	for _, arg := range frame.Args {
		if v, ok := arg.(string); ok {
			answer = v
			break
		}
	}

	res := ""
	if answer == p.data.GetGoal() {
		res = "홈런 !"
	} else {
		res = p.calculate(answer)
	}

	frame.Args = append(make([]any, 0), res)
	for _, session := range sessions {
		socket.Write(session.Conn, frame)
	}
}

// 클라이언트의 답안에서 스트라이크와 볼을 계산한다.
func (p *Process) calculate(answer string) string {
	strike, ball := 0, 0
	for i, x := range answer {
		index := strings.Index(p.data.GetGoal(), string(x))

		// 해당 숫자가 존재하지 않다면..
		if index == -1 {
			continue
		}

		// 존재하면, 자리 검사 후 자리도 같다면 strike, 아니면 ball을 증가한다.
		if index == i {
			strike++
		} else {
			ball++
		}
	}

	return fmt.Sprintf("%d 스트라이크, %d 볼", strike, ball)
}

var _ system.Process = new(Process)
