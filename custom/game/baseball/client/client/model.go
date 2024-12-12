package client

import (
	"fmt"
	"math/rand/v2"
	"software/import/client"
	"software/import/socket"
)

type Model struct {
	*client.Model
}

func New() *Model {
	return &Model{
		Model: client.New(
			&socket.Metadata{
				Name: fmt.Sprintf("%s%d", "플레이어", rand.Int()),
			},
		),
	}
}
