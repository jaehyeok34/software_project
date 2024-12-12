package socket

import (
	"encoding/json"
	"net"
)

// 클라이언트가 서버에 연결할 때, 메타데이터를 전송하는 과정이다.
func SendMetadata(conn net.Conn, meta *Metadata) error {
	return json.NewEncoder(conn).Encode(meta)
}

// 서버가 클라이언트의 연결을 허용할 때, 메타데이터를 수신하는 과정이다.
func ReceiveMetadata(conn net.Conn) (*Metadata, error) {
	meta := new(Metadata)
	if err := json.NewDecoder(conn).Decode(meta); err != nil {
		return nil, err
	}

	return meta, nil
}

// 프레임을 json 형태로 전달하는 과정이다.
func Write(conn net.Conn, frame *Frame) error {
	return json.NewEncoder(conn).Encode(frame)
}

// 프레임을 json 형태로 전달받는 과정이다.
func Read(conn net.Conn) (*Frame, error) {
	frame := new(Frame)
	if err := json.NewDecoder(conn).Decode(frame); err != nil {
		return nil, err
	}

	return frame, nil
}
