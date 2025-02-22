package main

import "fmt"

func any(f func(string) bool, a []string) bool {
	for _, i := range a {
		if f(i) == true {
			return true
		}
	}
	return false
}

func main() {
	a1 := []string{"Hello", "how", "are", "you"}
	a2 := []string{"This", "is", "4", "you"}

	result1 := any(IsNumeric, a1)
	result2 := any(IsNumeric, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}