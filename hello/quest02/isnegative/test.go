package main

import "fmt"

func isNegative(num int) {

	if num < 0 {
		fmt.Println("T")
	} else {
		fmt.Println("F")
	}
}

func main() {
	isNegative(1)
}
