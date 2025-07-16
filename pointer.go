package main

import "fmt"

func main() {
	age := 32

	var agePointer *int
	agePointer = &age

	getAdultAge(agePointer)
	fmt.Printf("Adult years: %d\n", age)
}

func getAdultAge(age *int) {
	*age = *age - 18
}
