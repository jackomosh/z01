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
	return answer + "\n"
}

func main() {
	fmt.Print(Firstword("hello how are you"))
	fmt.Print(Firstword("hello....    "))
	fmt.Print(Firstword("	hello....   there "))
	fmt.Print(Firstword(""))
}
