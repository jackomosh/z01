// to print uppercase alphabet from a to z
// to print lowercase alphabet from a to z

package main

import "fmt"

func main() {
	var PrintAlphabet int = 65

	for i := 0; i < 26; i++ {
		fmt.Print(string(rune(PrintAlphabet + i)))
	}

	// fmt.Println(string(rune(PrintAlphabet)))
}
