/**************************************************************
ANONYMOUS FUNCTIONIN GO
- No name
- Inline function
- Also known as `function literal`
- SYNTAX
	-- func(parameter_list)(return_type) {
		//code
	}()
**************************************************************/
package main

import (
	"fmt"
)

/*
func main() {
	func() {
		fmt.Println("This is ANONYMOUS FUNCTION")
	}()
}
*/

/**************************************************************
Anonymous function can be assigned to a variable.

func main() {
	value := func() {
		fmt.Println("Aonymous function is assigned to a variable!!!")
	}
	value()
}

**************************************************************/

/**************************************************************
Passing an anonymous function into another function as argument
**************************************************************/

func hello(f func(w2 string) string) {
	fmt.Println(f("Hello "))
}

func main() {
	world := func(w2 string) string {
		return w2 + "World"
	}
	hello(world)
}
