package main

import "fmt"

func PrintComb() {
	for i := 0; i <= 9; i++ {
		for j := i + 1; j <= 9; j++ {
			for k := j + 1; k <= 9; k++ {
				if i != 7 && j != 8 && k != 9 {
					fmt.Print(i)
					fmt.Print(j)
					fmt.Print(k)
					fmt.Print(",")
					fmt.Print(" ")
				}
			}
		}
	}
}

func main() {
	PrintComb()
}
