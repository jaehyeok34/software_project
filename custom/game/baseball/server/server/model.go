package server

import (
	"software/import/server"
)

type Model struct {
	*server.Model
	*BaseBall
}

func NewModel() *Model {
	return &Model{
		Model:    server.New(),
		BaseBall: NewBaseBall(),
	}
}
