package baseball

import (
	"software/import/server"
)

type Server struct {
	*server.Model
	data *Data
}

func NewServer(max int) *Server {
	return &Server{
		Model: server.New(),
		data:  NewData(max),
	}
}

// 현재 서버에 연결된 데이터를 반환한다.
func (s *Server) GetData() *Data {
	return s.data
}
