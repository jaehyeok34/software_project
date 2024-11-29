package socket

type Request struct {
	SystemType string        `json:"system_type"`
	Args       []interface{} `json:"args"`
}
