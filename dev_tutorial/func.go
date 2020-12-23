/**************************************************************
FUNCTION
- SYNTAX
	-- func (receiver)funcName(params) returnValues {
		//code
	}

- RECEIVER
	-- Connects function to a struct
**************************************************************/

package main

import (
	"fmt"
)

func greet() {
	fmt.Println("Hello, World!")
}

func greet_name(name string) string {
	return fmt.Sprintf("Name: %s", name)
}

func main() {
	greet()
	fmt.Println(greet_name("Harshal"))
}
