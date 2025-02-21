package main

import "fmt"

func Rot14(s string) string {
	str := []rune(s)  // Convert string to slice of runes to handle Unicode properly.

	for i := 0; i < len(str); i++ {
		// Handle lowercase letters ('a' to 'z')
		if str[i] >= 'a' && str[i] <= 'z' {
			// Shift by 14 positions and wrap around using modulo 26
			str[i] = ((str[i]-'a'+14)%26 + 'a')
		} else if str[i] >= 'A' && str[i] <= 'Z' { // Handle uppercase letters ('A' to 'Z')
			// Shift by 14 positions and wrap around using modulo 26
			str[i] = ((str[i]-'A'+14)%26 + 'A')
		}
	}
	return string(str)  // Convert rune slice back to string
}

func main() {
	result := Rot14("Hello! How are You?")

	fmt.Print(Rot14(result))
}
