/***************************************************************
POINTER
- Pointer is an addressing mechanism.
- `&` : addressing
- `*` : dereference
***************************************************************/

package main

import (
	"fmt"
)

func main() {
	a := 42
	m_a := &a
	fmt.Println("Value of a:", a)
	fmt.Println("Memory address of a:", &a)
	fmt.Println("Dereferencing memory location of a:", *m_a)
}
