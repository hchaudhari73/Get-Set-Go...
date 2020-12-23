package main

import (
	"fmt"
)

var (
	a = 10
	b = 2
)

func main() {
	d := a + b
	e := a - b
	f := a / b
	g := a * b
	h := a % b

	fmt.Printf("The result of %d + %d = %d\n", a, b, d)
	fmt.Printf("The result of %d - %d = %d\n", a, b, e)
	fmt.Printf("The result of %d / %d = %d\n", a, b, f)
	fmt.Printf("The result of %d * %d = %d\n", a, b, g)
	fmt.Printf("The result of %d modulo %d = %d\n", a, b, h)

}
