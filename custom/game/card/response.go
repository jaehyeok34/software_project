package card

import "fmt"

func (p *Player) receiveChat(message []interface{}) {
	for _, msg := range message {
		if str, ok := msg.(string); ok {
			fmt.Println("받은 메시지:", str)
		}
	}
}
