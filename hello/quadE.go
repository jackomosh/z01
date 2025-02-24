package main

import (
	"os"
	"strconv"
	"github.com/01-edu/z01"
)

func quadE(x, y int) {
	if x <= 0 || y <= 0 { // checks for positive integers
		return
	}
	for i := 0; i < y; i++ { // iterates over y-axis (h-height)
		for j := 0; j < x; j++ { // iterates over x-axis (w-width)
			if (i == 0 && j == 0) || (i == y-1 && j == x-1) {
				z01.PrintRune('A')
			} else if (i == 0 && j == x-1) || (i == y-1 && j == 0) {
				z01.PrintRune('C')
			} else if (i == 0 || i == y-1) || (j == 0 || j == x-1) {
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ') // prints a space if conditions are not met
			}
		}
		z01.PrintRune('\n')
	}
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	x, err1 := strconv.atoi(os.Args[1])
	y, err2 := strconv.atoi(os.Args[2])
	if err1 != nil || err2 != nil {
		return
	}
	quadE(x, y)
}
