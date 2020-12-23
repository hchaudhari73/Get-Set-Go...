package main

import (
	"fmt"
)

const pi = 3.142
const myconst string = "test constant"
const (
	a = 12
	b = 13
	c = 11
)

func main() {
	fmt.Println(pi)
	fmt.Println(a)
	// a = b // trying to change values of constant
	fmt.Println(a)
	fmt.Println(b)
	// b = c // trying to change values of constant
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%s, %T\n", myconst, myconst)
}

/*
# command-line-arguments
./const.go:18:4: cannot assign to a
./const.go:21:4: cannot assign to b
*/
