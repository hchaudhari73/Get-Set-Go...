/**************************************************************
FIRST CLASS CITIZEN IN GO
- The first class citizen principle in Go is also available but
rarely to use.
- The first class citizen means that a function can be used as
a argument or parameter in a function.
**************************************************************/

package main

import (
	"fmt"
)

func foo(f func(xi ...int) int, data ...int) int {
	return f(data...)
}

func main() {
	//declare a dunction that stored inside variable called cb
	cb := func(xi ...int) int {
		total := 0
		for _, value := range xi {
			total += value
		}
		return total
	}

	//declare a slice in int
	data := []int{1, 2, 3, 4, 5}
	//call foo() function with arguments:
	//1. cb means callbac function
	//2. data... means retrieve all items from slice called
	//	"data"
	result := foo(cb, data...)
	fmt.Println(result)
}
