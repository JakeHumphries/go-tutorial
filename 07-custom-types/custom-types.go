package main

import "fmt"

type hotdog int
var b hotdog 

func main() {
	b = 45
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
