package socket

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func Send(conn net.Conn, data []byte) error {
	length := len(data)
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(length))

	if _, err := conn.Write(append(buf, data...)); err != nil {
		fmt.Println("sendData 문제발생:", err)
		return err
	}

	fmt.Println("데이터 보냄")
	return nil
}

func Receive(conn net.Conn) ([]byte, error) {
	buf := make([]byte, 4)
	if _, err := io.ReadFull(conn, buf); err != nil {
		fmt.Println("read 문제 발생(header):", err)
		return nil, err
	}

	length := binary.BigEndian.Uint32(buf)
	buf = make([]byte, length)
	if _, err := io.ReadFull(conn, buf); err != nil {
		fmt.Println("read 문제 발생(payload):", err)
		return nil, err
	}

	fmt.Println("received message:", string(buf))
	return buf, nil
}
