package baseball

import (
	"math/rand/v2"
	"slices"
	"strconv"
)

type Data struct {
	goal    []int
	isStart bool
	count   int
}

func NewData() *Data {
	return new(Data)
}

// 게임이 시작되면, 목표 숫자와 시작 여부, 답안 제출 횟수를 초기화한다.
func (d *Data) Init() bool {
	// 게임이 이미 시작된 상태(초기화 완료)이면, 초기화를 무시한다.
	if d.isStart {
		return false
	}

	d.goal = make([]int, 0)
	d.isStart = true
	d.count = 0

	// 세 자리의 숫자를 생성한다.
	// 생성된 랜덤 값이 중복된 값이 있다면, 재생성한다.
	for len(d.goal) < 3 {
		n := rand.IntN(10)
		if !slices.Contains(d.goal, n) {
			d.goal = append(d.goal, n)
		}
	}

	return true
}

// []int 형태로 저장되어 있는 goal을 하나의 문자열 "XXX"로 반환한다.
func (d *Data) GetGoal() string {
	goal := ""
	for _, v := range d.goal {
		goal += strconv.Itoa(v)
	}

	return goal
}
