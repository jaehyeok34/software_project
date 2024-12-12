package main

import (
	"fmt"
	"software/import/collection"
)

type A struct {
	Value string
}

func main() {
	m := collection.NewMap[string, *A]()
	a := &A{Value: "hello"}

	m.Store(a.Value, a)

	if v, ok := m.Load("hello"); ok {
		fmt.Println(v)
	}

	fmt.Println(m.Length())
}
