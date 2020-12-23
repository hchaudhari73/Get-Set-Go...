/**************************************************************
INTERFACE
- Similar with `struct` but it only contains some abstract
methods.
- SYNTAX
	-- declare a `struct`
	-- declare a `interface`
	-- declare a `function` to pass in `interface`
- METHOD SETS
	-- method set is a fucntion that associates with certain
	type.
	e.g. area() associates with struct shape.
	-- This `area()` method is available for value and pointer
	of type shape.
	-- But if the receiver of method set is a pointer, then the
	method is only available with pointer.
		-- eg func (r *rectangle) area() int {
			return r.length * r.width
		}
**************************************************************/

package main

import (
	"fmt"
)

type rectangle struct {
	length int
	width  int
}

type shape interface {
	area() int
}

type parameter interface {
	param() int
}

type sides interface {
	get_sides() [2]int
}

func (r rectangle) area() int {
	return r.length * r.width
}

func (r rectangle) param() int {
	return 2 * (r.length + r.width)
}

//receiver of the method set is a pointer
//to call this we have to always give the pointer
func (r *rectangle) get_sides() [2]int {
	side := [2]int{r.length, r.width}
	return side
}

func info(s shape, p parameter, si sides) {
	fmt.Printf("Area of rectangle is %d.\n", s.area())
	fmt.Printf("Parameter of rectangle is %d.\n", p.param())
	fmt.Printf("length %d, width %d\n", si.get_sides()[0], si.get_sides()[0])
}

func main() {
	r := rectangle{
		length: 10,
		width:  12,
	}
	info(r, &r, &r)
}
