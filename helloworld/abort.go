package main 

import "fmt"

func abort(a,b,c,d,e int) int {
	answer := [5]int{a,b,c,d,e}
	mid := len(answer)/2
	
	for i := 0; i < len(answer) - 1; i++ {
		for j := 0; j < len(answer) - 1 - i; j++ {
			if answer[j] > answer[j+1] {
				answer[j], answer[j+1] = answer[j+1], answer[j]
			}
		}
	}
	return answer[mid]
}

func main() {
	middle := abort(2, 3, 8, 5, 7)
	fmt.Println(middle)
}