package main

import "fmt"

func fishAndchips(n int) string {
	if n < 0 {
		return ("error: number is negative")
	} else if n%2 == 0 && n%3 == 0 {
		fmt.Print("\nfish and chips")
	} else if n%2 == 0 {
		fmt.Print("fish")
	} else if n%3 == 0 {
		fmt.Print("\nchips")
	} else {
		return ("error: non divisible")
	}
	return string(n)
}

func main() {
	fishAndchips(4)
	fishAndchips(9)
	fishAndchips(6)
	// fishAndchips(20)
}
