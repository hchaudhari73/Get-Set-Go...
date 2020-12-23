/**************************************************************
ANONYMOUS STRUCT
- SYNTAX
	-- s1 := struct {
		//declare some fields
		fields int
		fields []string
	}{
		//instantiate directly
		field : 12,
		field : []string{"hi", "mate"},
	}
**************************************************************/

package main

import (
	"fmt"
)

func main() {
	person := struct {
		name string
		age  int
	}{
		name: "Harshal",
		age:  24,
	}
	fmt.Println(person.name)
	fmt.Println(person.age)
}
