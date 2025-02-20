package main

import "fmt"

func containsAlphabet(str string) bool {
	for _, char := range str {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') {
			return false
		}
	}
	return true
}

func isUpper(a rune) bool {
	return a >= 'A' && a <= 'Z'
}

func camelCase(s string) string {
	results := ""

	if len(s) == 0 || !containsAlphabet(s) {
		return s
	}
	for i := 0; i < len(s); i++ {
		if i != 0 && isUpper(rune(s[i])) && i+1 < len(s) && !isUpper(rune(s[i+1])) {
			results += "_"
			results += string(s[i])
		} else if !isUpper(rune(s[i])) || (i == 0 && isUpper(rune(s[i]))) {
			results += string(s[i])
		} else {
			return s
		}
	}
	return results
}

func main() {
	s := "helloWorld"
	fmt.Println(camelCase(s))
}
