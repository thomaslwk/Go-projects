package main

import (
	"fmt"
	"strconv"
)

// Create intlist slice of type string
type intlist []string

func main() {
	// Initialise intlist
	intlist1 := intlist{}

	// For loop to iterate 10 times.
	for i := 1; i < 11; i++ {
		// Convert int val. to string
		intval := strconv.Itoa(i)
		// Even, ODD check
		if i%2 == 0 {
			intlist1 = append(intlist1, intval+" is even")
		} else {
			intlist1 = append(intlist1, intval+" is odd")
		}
	}
	// Print function
	intlist1.print()
}

func (intl intlist) print() {
	for _, intval := range intl {
		fmt.Println(intval)
	}
}
