package main

import "fmt"

func main() {
	str := "Go is Awesome"

	var result []rune

	for _, char := range str {
		if char != ' ' {
			result = append(result, char)
		}
	}
	fmt.Println(string(result))
}
