package main

import "fmt"

func main() {
	ninja1()
	ninja2()
	ninja3()
	ninja4()
	ninja5()
}

// Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”
// 42
// “James Bond”
// true
// Now print the values stored in those variables using
// a single print statement
// multiple print statements

func ninja1() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Printf("%v\t%v\t%v\n", x, y, z)
	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}

// Use var to DECLARE three VARIABLES. The variables should have package level scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that TYPE).
// identifier “x” type int
// identifier “y” type string
// identifier “z” type bool
// in func main
// print out the values for each identifier
// The compiler assigned values to the variables. What are these values called?

var a int
var b string
var c bool

func ninja2() {
	println(a, b, c)
	// these are called zero values
}

// At the package level scope, assign the following values to the three variables
// for x assign 42
// for y assign “James Bond”
// for z assign true
// in func main
// use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
// print out the value stored by variable “s”

var x = 42
var y = "James Bond"
var z = true

func ninja3() {
	s := fmt.Sprintf("%v\t%v\t%v\n", x, y, z)
	fmt.Println(s)
}

// Create your own type. Have the underlying type be an int.
// create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
// in func main
// print out the value of the variable “x”
// print out the type of the variable “x”
// assign 42 to the VARIABLE “x” using the “=” OPERATOR
// print out the value of the variable “x”

type customInt int

var q customInt

func ninja4() {
	fmt.Println(q)
	fmt.Printf("%T\n", q)
	q = 42
	fmt.Println(q)

}

// at the package level scope, using the “var” keyword, create a VARIABLE with the IDENTIFIER “y”. The variable should be of the UNDERLYING TYPE of your custom TYPE “x”
// in func main
// this should already be done
// print out the value of the variable “x”
// print out the type of the variable “x”
// assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
// print out the value of the variable “x”
// now do this
// now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
// then use the “=” operator to ASSIGN that value to “y”
// print out the value stored in “y”
// print out the type of “y”

var w int

func ninja5() {
	fmt.Println(q)
	fmt.Printf("%T\n", q)
	fmt.Println(q)
	w = int(x)
	fmt.Println(w)
	fmt.Printf("%T\n", w)
}
