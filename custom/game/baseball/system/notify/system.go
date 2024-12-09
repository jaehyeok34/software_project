package notify

import (
	"fmt"
	"net"
	"software/custom/game/baseball/model"
	"software/import/room"
	"software/import/socket"
)

type System struct {
	model *model.Model
}

func (s *System) Run(src *room.Connection, conns []net.Conn, args ...interface{}) {
	// 순서에 해당하는 플레이어에게 메시지를 보내는 로직
	var message string
	for _, arg := range args {
		if v, ok := arg.(string); ok {
			message = v
		}
	}

	// model에 있는 goal과 message를 비교하는 과정
	strike, ball := s.model.Compare(message)
	data := &socket.Frame{}
	data.Name = src.Name
	data.Event = "notify"

	if strike == 3 {
		data.Args = append(data.Args, "홈런!!")
	} else {
		data.Args = append(data.Args, fmt.Sprintf("%d 스트라이크, %d 볼", strike, ball))
	}

	for _, conn := range conns {
		socket.Write(conn, data)
	}
}
