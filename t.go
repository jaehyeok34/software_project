package main

import "fmt"

type T interface {
	Run()
}

func TT() any {
	m := make(map[string]T)

	return m["hello"]
}

func main() {
	t, ok := TT().(T)
	if !ok {
		fmt.Println("???")
	}

	fmt.Println(t)
}
