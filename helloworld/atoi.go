package main

import "fmt"

// func Atoi(s string) int {
// 	var n int
// 	sliceS := []rune(s)
// 	first := true
// 	for i := 0; i < len(sliceS); i++ {
// 		if (sliceS[0] == '-' || sliceS[0] == '+') && first {
// 			first = false
// 			continue
// 		}
// 		if sliceS[i] > '9' || sliceS[i] < '0' {
// 			return 0
// 		}
// 		n = n*10 + int(sliceS[i]-'0')
// 	}
// 	if sliceS[0] == '-' {
// 		return -n
// 	}
// 	return n
// }

func Atoi(s string) int {
	var n int
	slices := []rune(s)
	first := true

	for i := 0; i < len(slices); i++ {
		if slices[0] == '-' || slices[0] == '+' && first {
			first = false
			continue
		}
		if slices[i] > '9' || slices[i] < '0' {
			return 0
		}
		n = n*10 + int(slices[i]-'0')
	}
	if slices[0] == '-' {
		return -n
	}
	return n
}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
}

// quest 3 / checkpoint
