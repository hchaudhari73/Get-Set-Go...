/***************************************************************
PANIC
- `panic` function is error handling mechanism in Go.
- SYNTAX
	-- panic("cannot divide by 0")
***************************************************************/

package main

import (
	"fmt"
)

func divide(x, y int) int {
	if y == 0 {
		panic("cannot divide by 0")
	} else {
		return x / y
	}
}

func main() {
	fmt.Println(divide(5, 2))
}
