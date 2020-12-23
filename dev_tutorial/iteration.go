/**************************************************************
ITERATION
	-- for (init condition; condition; post condition) { ... }
**************************************************************/

package main

import (
	"fmt"
)

/*
func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}
*/

/*
func main() {
	i := 1
	for {
		fmt.Println(i)
		if i >= 10 {
			break
		}
		i++
	}
}
*/

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println()
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
	}
}
