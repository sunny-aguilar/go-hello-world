package main

import (
	"fmt"
	"time"
)

import "rsc.io/quote"

func main() {
	// variables
	var i int = 42 // explicit variable
	j := 42        // compiler inferred variable
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(quote.Go())

	var n bool = true
	fmt.Printf("%v, %T", n, n)

	now := time.Now()
	fmt.Println(now)

	var roomNumber, floorNumber int = 154, 3
	fmt.Println(roomNumber, floorNumber)

	var password = "notSecured"
	fmt.Println(password)

	// constants
	const neverChange int = 100
	fmt.Println(neverChange)

	// control statements
	if 1 < 2 {
		fmt.Printf("2 is greater than 1\n")
	}

	// for loops
	var ctr int = 0
	for ctr < 2 {
		fmt.Printf("I is less than 2, %v\n", ctr)
		ctr++
	}

	for i:= 0; i < 3; i++ {
		fmt.Printf("i=%v\n", i)
	}
}

// go run . - runs a Go program
// go build . - creates executable that can be run later
// ./GoExecutable - runs the Go executable
