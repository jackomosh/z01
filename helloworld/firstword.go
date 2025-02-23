package main

import "fmt"

func Firstword(s string) string {
	var answer string

	for i := range s {
		if s[i] == ' ' || s[i] == '\t' || s[i] == '.' {
			answer = s[:i]
			break
		}
	}
	return answer
}

func main() {
	fmt.Println(Firstword("hello how are you"))
	fmt.Println(Firstword("hello....    "))
	fmt.Println(Firstword("	hello....   there "))
}
