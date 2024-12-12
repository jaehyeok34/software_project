package start

import (
	baseball "software/custom/game/baseball/server/server"
	"software/import/server"
	"software/import/socket"
)

type System struct {
	baseball *baseball.BaseBall
}

func New(baseball *baseball.BaseBall) *System {
	return &System{baseball: baseball}
}

var Event = "start"

func (s *System) Run(src *server.Session, dsts []*server.Session, frame *socket.Frame) {
	var msg string
	if s.baseball.Init() {
		msg = s.baseball.GetGoal()
	} else {
		msg = "이미 게임이 시작됐습니다."
	}

	frame.Meta = nil
	frame.Args = []any{msg}
	for _, dst := range dsts {
		socket.Write(dst.Conn, frame)
	}
}

var _ server.System = new(System)
