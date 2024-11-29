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
	req := new(Frame)
	decoder := json.NewDecoder(conn)
	if err := decoder.Decode(req); err != nil {
		return nil, err
	}

	return req, nil
}

// --------------길이 기반 통신인데 json 아니라서 일단 주석--------------------
// func Send(conn net.Conn, data []byte) error {
// 	length := len(data)
// 	buf := make([]byte, 4)
// 	binary.BigEndian.PutUint32(buf, uint32(length))

// 	if _, err := conn.Write(append(buf, data...)); err != nil {
// 		fmt.Println("sendData 문제발생:", err)
// 		return err
// 	}

// 	fmt.Println("데이터 보냄")
// 	return nil
// }

// func Receive(conn net.Conn) ([]byte, error) {
// 	buf := make([]byte, 4)
// 	if _, err := io.ReadFull(conn, buf); err != nil {
// 		fmt.Println("read 문제 발생(header):", err)
// 		return nil, err
// 	}

// 	length := binary.BigEndian.Uint32(buf)
// 	buf = make([]byte, length)
// 	if _, err := io.ReadFull(conn, buf); err != nil {
// 		fmt.Println("read 문제 발생(payload):", err)
// 		return nil, err
// 	}

// 	fmt.Println("received message:", string(buf))
// 	return buf, nil
// }
