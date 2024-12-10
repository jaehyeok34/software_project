package socket

type Frame struct {
	Meta  *Metadata     `json:"meta"`
	Event string        `json:"event"`
	Args  []interface{} `json:"args"`
}
