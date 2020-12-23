/***************************************************************
DEFER
- There is a case when function neds to be executed at certain
time for eg. funcA() has to be executed after certain func is
finished.
- `defer` basically pushes a function into a list. List of
functions is executed after the surroundings function's execution
is finished.
- SYNTAX
	-- defer func() {....}
***************************************************************/

package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

func main() {
	x := 5
	y := 10
	defer fmt.Println("FIRST")
	fmt.Println(add(x, y))
}

/*
15
FIRST
*/
