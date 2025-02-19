package main

import "fmt"

func IsMedian() int {

	answer := []int{10, 50, 80, 60, 70, 40, 30}
	mid := len(answer)/2

	if len(answer) % 2 != 0 {
		return answer[mid]
	}
	return answer[mid]
}

func main() {
	fmt.Print(IsMedian())
}
