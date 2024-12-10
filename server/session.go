package server

import (
	"net"
	"software/socket"
)

type Session struct {
	Meta *socket.Metadata `json:"meta"`
	Conn net.Conn         `json:"conn"`
}
