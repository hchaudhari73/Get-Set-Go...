/***************************************************************
RECURSION
- A function that calls itself.
***************************************************************/
package main

import "fmt"

//Factorial

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}

func main() {
	NUMBER := 5
	fmt.Println(factorial(NUMBER))
}
