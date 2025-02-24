package main

import "fmt"

func alphaCount(str string) int {
	count := 0
	for _, c := range str {
		if c >= 'A' && c <= 'Z' || c >= 'a' && c <= 'z' {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(alphaCount("hello2    "))
}
