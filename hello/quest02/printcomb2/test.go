package main

import "fmt"

func PrintComb() {
	for i := 0; i <= 10; i++ {
		for j := i + 1; j <= 9; j++ {
			if i != 10 || j != 9 {
				fmt.Print(i, j)
				fmt.Print(", ")
			}
		}
	}
}

func main() {
	PrintComb()
}
