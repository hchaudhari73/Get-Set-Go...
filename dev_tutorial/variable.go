package main

import (
	"fmt"
)

var (
	a = 1
	b = 2
	c = 3
) // declare multiple variables

func main() {
	data := "Box this lap" //declare variable with shorthand syntax
	fmt.Println(a, b, c)
	fmt.Printf("%T %T %T\n", a, b, c) //print the type of a, b, c variable with %T notation
	fmt.Println(data)
	fmt.Printf("%s %T\n", data, data)
}
