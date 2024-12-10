package main

import (
	"fmt"
	"time"
)

type Test struct {
	List map[string]int
}

func main() {
	t := new(Test)
	t.List = make(map[string]int)

	t.List["key"] = 100

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i, ":", t.List["key"])
			time.Sleep(time.Millisecond * 30)
		}
	}()

	time.Sleep(time.Millisecond * 500)
	fmt.Println("삭제")
	delete(t.List, "key")

	for {
	}
}
