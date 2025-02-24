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
	return answer + "\n"
}

func main() {
	fmt.Print(lastWord("hello how are you"))
	fmt.Print(lastWord("hello....    "))
	fmt.Print(lastWord("	hello....   there."))
	fmt.Print(lastWord(""))
}
