package main

import "fmt"

func CountRunes(s string) int {
	return len([]rune(s))
}

func main() {
	str := "Hello world!"

	fmt.Println("RuneCount:", CountRunes(str))
}
