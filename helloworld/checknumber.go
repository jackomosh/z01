package main 

import "fmt"

func checkNumber(str string) bool {
	for _, char := range str {
		if char >= '0' && char <= '9' {
			return true
		}
	}
	return false
}

func main() {
	a := "hell1o"
	fmt.Println(checkNumber(a))
}