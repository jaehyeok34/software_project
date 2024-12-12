package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "12"

	fmt.Println(utf8.RuneCountInString(s))
}
