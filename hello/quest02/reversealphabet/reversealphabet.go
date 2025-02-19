// to print uppercase alphabet from z to a
// to print lowercase alphabet from z to a

package main

import "fmt"

func main() {
	var reversealphabet int = 122

	for i := 0; i < 26; i++ {
		fmt.Print(string(rune(reversealphabet - i)))
	}
}
