package socket

type Frame struct {
	Event string        `json:"event"`
	Args  []interface{} `json:"args"`
}
