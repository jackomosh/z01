package main

import (
	"os"
	"strconv"
	"github.com/01-edu/z01"
)

func quadD(x, y int) {
    if x <= 0 || y <= 0 { // checks for positive integers
        return
    }
    for h := 0; h < y; h++ { // iterates over y-axis (h-height)
        for w := 0; w < x; w++ { // iterates over x-axis (w-width)
            if (h == 0 && w == 0) || (h == y-1 && w == 0) {
                z01.PrintRune('A')
            } else if (h == 0 && w == x-1) || (h == y-1 && w == x-1) {
                z01.PrintRune('C')
            } else if (h == 0 || h == y-1) || (w == 0 || w == x-1) {
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
	x, err1 := strconv.Atoi(os.Args[1])
	y, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil {
		return
	}
	quadD(x, y)
}