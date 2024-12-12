package main

import "fmt"

func main() {
	var a any = 10

	fmt.Println(a.(string))
}
