package main

import(
	"fmt"
	"strconv"
)

func main() {
	fmt.Print(Zipstring("YouuungFelllla"))
	fmt.Print(Zipstring("Helloo Therre"))
}

func Zipstring (s string) string {
	var r string
	var c int

	for i := 0; i < len(s); i+=c {
		c = 0
		for j := i; j < len(s); j++ {
			if s[j] != s[i] {
				break
			}
			c++
		}
		r += Itoa(c) + string(s[i])
	}
	return r  + "\n"
}

func Itoa(n int) (s string) {
	return strconv.Itoa(n)
}
