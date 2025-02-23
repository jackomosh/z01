package main

import "fmt"

func lastWord(s string) string {
	var answer string

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' || s[i] == '\t' {
			answer = s[i+1:]
			break
		}
	}
	return answer
}

func main() {
	fmt.Println(lastWord("hello how are you"))
	fmt.Println(lastWord("hello....    "))
	fmt.Println(lastWord("	hello....   there."))
}
