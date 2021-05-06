package main

import "fmt"

// Declare and assign variable
// Can be declared outside a function
var y int = 43
var z int

func main() {
	// Declare and assign variable at the same time
	// can only be used inside a function
	x := 42
	fmt.Println(x)

	fmt.Println(z)
	foo()
}

func foo() {
	fmt.Println(y)
}

// Limit your scope!
// Y's scope is package level
// Z's value is 0 as we didnt assign a value