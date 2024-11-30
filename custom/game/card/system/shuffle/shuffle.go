package shuffle

import (
	"fmt"
	"net"
	"software/import/room"
	"software/import/socket"
)

type System struct{}

var Key string = "shuffle"

func New() *System {
	return &System{}
}

func (s *System) Run(src *room.Connection, conns []net.Conn, args ...interface{}) {
	if len(conns) == 0 {
		fmt.Println("ChatSystem(Run): missing []net.Conn")
		return
	}

	// ... 카드를 섞는 알고리즘

	fmt.Println("카드를 섞는 중 입니다.")
	fmt.Println("카드 섞기가 완료했습니다.")

	f := new(socket.Frame)
	f.Name = src.Name
	f.Event = Key
	f.Args = append(f.Args, "카드를 섞었습니다.")
	socket.Write(src.Conn, f)
}

var _ room.System = new(System)
