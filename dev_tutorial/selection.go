/**************************************************************
SELECTION
- selection is a mechanism to execute certain code based on
specified condition.package dev_tutorial
	-- if
	-- if else
	-- if else if else
	-- switch case

OPERATORS
	-- || OR
	-- && AND
	-- > greater than
	-- >= greater than equal to
	-- <= less than equal to
	-- < less than
	-- == is equals?
	-- != is not equals?
	-- ! negation of statement
**************************************************************/
package main

/*
func main() {
	a := 42
	if a%2 == 0 {
		fmt.Printf("%d is even number\n", a)
	} else {
		fmt.Printf("%d is odd number\n", a)
	}
}
*/

/*
func main() {
	user := ""
	switch user {
	case "user":
		fmt.Println("welcome user!")
	case "admin":
		fmt.Println("welcome admin!")
	default:
		fmt.Println("Who are you?")
	}
}
*/

/**************************************************************
There is `fallthrough` keyword in switch case selection to
execute another code under the matched case.
**************************************************************/

func main() {
	user := "user"
	switch user {
	case "user":
		println("Welcome user")
		fallthrough
	case "admin":
		println("welcome admin")
	case "manager":
		println("welcome manager")
	}
}
