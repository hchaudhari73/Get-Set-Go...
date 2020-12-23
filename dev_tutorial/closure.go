/***************************************************************
CLOSURE
- Closure means a fucntion that could access variables form
outside it body.
***************************************************************/
package main

import "fmt"

func increment() func() int {
	i := 0 //forms a closure
	return func() int {
		i++
		return i
	}
}

func main() {
	myInt := increment()
	fmt.Println(myInt())
	fmt.Println(myInt())
	fmt.Println(myInt())
}
