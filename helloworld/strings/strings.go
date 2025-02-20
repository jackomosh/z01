package main

import "fmt"

func main() {
	name := "Jack Omondi"
	fmt.Printf("My name is: %v", name) // prints name to stdout
	fmt.Printf("\nLast index: %v", len(name)-1) // prints the last index but in byte form

	name2 := name[0]
	name3 := name[len(name)-1]
	firstName := name[:4]
	LastName := name[5:]
	thirdfour := name[3:6]
	fmt.Print("\nLetter at index 0 is: ", string(name2)) // prints the character of name at index 0
	fmt.Print("\nLetter at last index is: ", string(name3)) // prints the character of name at last index
	fmt.Print("\nYour Firstname is: ", firstName) // prints the first words of name
	fmt.Print("\nYour Lastname  is: ", LastName) // prints the first words of name
	fmt.Print("\nYour Lastname  is: ", thirdfour)
}