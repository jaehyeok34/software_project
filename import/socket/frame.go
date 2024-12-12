package socket

// 소켓 통신에 사용되는 실제 데이터 단위이다.
type Frame struct {
	Meta  *Metadata `json:"meta"`
	Event string    `json:"event"`
	Args  []any     `json:"args"`
}
