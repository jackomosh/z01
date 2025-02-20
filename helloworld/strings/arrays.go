package main 

import "fmt"

func main() {
	var ages = []int{20,30,40}

	var names = []string{"Jack", "James", "Leon", "Olivia"}
	names[3] = "John"

	ages = append(ages, 50)

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))
}