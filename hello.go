package main

import "fmt"

import "rsc.io/quote"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
}

// go run . - runs a Go program
// go build . - creates executable that can be run later
// ./GoExecutable - runs the Go executable
