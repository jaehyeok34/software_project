package server

import (
	"fmt"
	"math/rand/v2"
	"slices"
	"strconv"
)

type BaseBall struct {
	goal    []int
	isStart bool
	count   int
}

func NewBaseBall() *BaseBall {
	return &BaseBall{make([]int, 0), false, 0}
}

func (b *BaseBall) Init() bool {
	if b.isStart {
		return false
	}

	b.goal = make([]int, 0)
	b.isStart = true
	b.count = 0

	for len(b.goal) < 3 {
		n := rand.IntN(10)

		if !slices.Contains(b.goal, n) {
			b.goal = append(b.goal, n)
		}
	}

	fmt.Println(b.goal)
	return true
}

func (b *BaseBall) GetGoal() string {
	s := ""
	for _, v := range b.goal {
		s += strconv.Itoa(v)
	}

	return s
}
