/***************************************************************
Testing function
***************************************************************/

package main

import (
	"fmt"
)

func add_two(x int) int {
	result := x + 2
	return result
}

func main() {
	result := add_two(4)
	fmt.Println("Sum =", result)
}
