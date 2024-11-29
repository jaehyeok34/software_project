package socket

import (
	"encoding/json"
	"fmt"
	"net"
)

func Write(conn net.Conn, data *Frame) error {
	encoder := json.NewEncoder(conn)
	if err := encoder.Encode(data); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func Read(conn net.Conn) (*Frame, error) {
	f := new(Frame)
	decoder := json.NewDecoder(conn)
	if err := decoder.Decode(f); err != nil {
		return nil, err
	}

	return f, nil
}
