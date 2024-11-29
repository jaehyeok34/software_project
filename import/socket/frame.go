package socket

type Frame struct {
	Name  string        `json:"name"`
	Event string        `json:"event"`
	Args  []interface{} `json:"args"`
}
