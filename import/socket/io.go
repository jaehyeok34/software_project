package socket

import (
	"encoding/json"
	"net"
)

func RequestConnection(conn net.Conn, meta *Metadata) error {
	return json.NewEncoder(conn).Encode(meta)
}

func RecieveConnection(conn net.Conn) (*Metadata, error) {
	meta := new(Metadata)
	if err := json.NewDecoder(conn).Decode(meta); err != nil {
		return nil, err
	}

	return meta, nil
}

func Write(conn net.Conn, frame *Frame) error {
	return json.NewEncoder(conn).Encode(frame)
}

func Read(conn net.Conn) (*Frame, error) {
	frame := new(Frame)
	if err := json.NewDecoder(conn).Decode(frame); err != nil {
		return nil, err
	}

	return frame, nil
}
