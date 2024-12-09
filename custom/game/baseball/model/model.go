package model

import (
	"fmt"
	"math/rand"
	"software/import/room"
)

type Model struct {
	room *room.Model
	goal string
}

func New() *Model {
	return &Model{room: room.New()}
}

func (m *Model) Init() {
	first := rand.Intn(10)

	var second int
	for {
		second = rand.Intn(10)
		if first != second {
			break
		}
	}

	var third int
	for {
		third = rand.Intn(10)
		if first != third && second != third {
			break
		}
	}

	m.goal = fmt.Sprintf("%d%d%d", first, second, third)
	fmt.Println(m.goal)
}

func (m *Model) Listen(network string, address string, n uint) error {
	return m.room.Listen(network, address, n)
}

func (m *Model) AddSystem(key string, system room.System) {
	m.room.AddSystem(key, system)
}

func (m *Model) Compare(answer string) (int, int) {
	strike := 0
	ball := 0

	for i := 0; i < len(answer); i++ {
		if m.contains(answer[i]) {
			if answer[i] == m.goal[i] {
				strike++
				continue
			}

			ball++
		}
	}

	return strike, ball
}

func (m *Model) contains(char byte) bool {
	for i := 0; i < len(m.goal); i++ {
		if m.goal[i] == char {
			return true
		}
	}
	return false
}
