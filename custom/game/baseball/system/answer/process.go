package answer

import (
	"fmt"
	"software/custom/game/baseball"
	"software/import/socket"
	"software/import/system"
	"strconv"
	"strings"
	"unicode/utf8"
)

var Event = "answer"

type process struct {
	data *baseball.Data
}

func NewProcess(data *baseball.Data) system.Process {
	return &process{data}
}

// 클라이언트로부터 답안 제출 이벤트를 수신했을 때 처리하는 로직이다.
func (p *process) Run(src *socket.Metadata, frame *socket.Frame, sessions []*socket.Session) {
	r := ""
	if !p.data.IsStart {
		r = "게임을 먼저 시작해 주세요."
	} else {
		r = p.getResulst(frame.Args)
	}

	frame.Args = append(make([]any, 0), r)
	for _, session := range sessions {
		socket.Write(session.Conn, frame)
	}
}

// 수신한 프레임으로부터 답안을 추출하고, 결과를 반환한다.
func (p *process) getResulst(args []any) string {
	answer := ""
	for _, arg := range args {
		if v, ok := arg.(string); ok {
			answer = v
			break
		}
	}

	// 수신한 답이 세자리 문자가 아니거나 숫자가 아닐 경우에는 잘못된 입력으로 판단한다.
	_, err := strconv.Atoi(answer)
	if utf8.RuneCountInString(answer) != 3 || err != nil {
		return "잘못된 입력입니다 !"
	}

	// 기회를 감소했을 때, 0이라면 게임을 종료한다.
	p.data.Count--
	if p.data.Count == 0 {
		p.data.IsStart = false
		return "게임 오버 !"
	}

	// 답을 맞췄을 경우 게임을 종료한다.
	if answer == p.data.GetGoal() {
		p.data.IsStart = false
		return "홈런 !"
	}

	// n 스트라이크, m 볼의 결과와 남은 기회를 반환한다.
	return fmt.Sprintf(
		"%s\n%s",
		p.calculate(answer),
		fmt.Sprintf("남은 기회는 %d번 입니다.", p.data.Count),
	)
}

// 클라이언트의 답안에서 스트라이크와 볼을 계산한다.
func (p *process) calculate(answer string) string {
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

var _ system.Process = new(process)
