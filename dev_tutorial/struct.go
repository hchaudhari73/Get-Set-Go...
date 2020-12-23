/**************************************************************
FUNCTION WITH RECEIVER
- SYNTAX
	-- struct
	-- type person struct {
		name string
	}
	   func (p person) greet() {
		   fmt.Printf("Hello %s", p.name)
	   }

ANONYMOUS FUNCTION
- Function without name or identifier.
- IIFE : Immediately-invoked Function Expression
		We declare a function that executes directly
- SYNTAX
	-- func() {
		//code
	   } //using IIFE
	-- func() {
		//code
	}() //use this notation to execute the function
**************************************************************/

package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

// declare a method say() with type person as receiver
func (p person) say() {
	fmt.Println("Hello, my name is:", p.name)
}

func main() {

	p := person{
		name: "Harshal",
		age:  24,
	}

	fmt.Printf("Name %s\n", p.name)
	fmt.Printf("Age %d\n", p.age)
	p.say()
}

/*
//anonymous function
func() {
	//calculates sum
	marks := []int{1, 2, 3, 4, 5}
	total := 0
	for _, value := range marks {
		total += value
	}
	fmt.Println(total)
}()
*/
