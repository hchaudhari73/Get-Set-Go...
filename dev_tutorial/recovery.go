/***************************************************************
RECOVERY
- To recover a program that is panicking.
- `recover` function ca be used in deferred function.
- `recover` func is useless if used in non deferred func, because
if the prog is panicking, the deferred func still executed.
- SYNTAX
	-- defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered!")
		}
	}()
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
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered!")
		}
	}()
	fmt.Println(divide(5, 0))
}
