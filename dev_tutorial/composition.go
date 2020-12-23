/**************************************************************
COMPOSITION
- Substitute of inheritance.
- SYNTAX
	-- type s1 struct{
		name string
		age string
	}
	-- type s2 struct{
		s1
		team string
	}
	-- p = s1{...}
	-- p2 = s2{p, ....,}
**************************************************************/
package main

import (
	"fmt"
)

//declare a struct called person
type person struct {
	name string
	age  int
}

//composition example
//declare a struct called manager that includes another struct
type manager struct {
	person
	team string
}

func main() {
	p := person{
		name: "Harshal",
		age:  24,
	}
	man := manager{
		person: p,
		team:   "Development",
	}
	fmt.Printf("Hello, my name is %s, and my age is %d, and I am from %s, team",
		man.name, man.age, man.team)
}
