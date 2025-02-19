package main

import "fmt"

func QuadA(x, y int) {
	if x <= 0 &&  y <= 0 {
		return
	}

	for h := 0; h < y; h++ {
		for w := 0; w < x; w++ {
			if h == 0 && w == 0 {
				fmt.Print("o")
			} else if h == 0 && w == x - 1 {
				fmt.Print("o")
			} else if h == y - 1 && w == 0 {
				fmt.Print("o")
			} else if h == y - 1 && w == x - 1 {
				fmt.Print("o")
			} else if h == 0 || h == y - 1 {
				fmt.Print("-")
			} else if w == 0 || w == x - 1 {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	QuadA(5, 3)
}