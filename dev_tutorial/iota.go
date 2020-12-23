/**************************************************************
iota is special const in go.
it's default value = 0 and will increased automatically by 1
if assigned into another constant.
***************************************************************/
package main

import (
	"fmt"
)

const (
	a = iota
	b = iota
	c = iota
)

// second example

const (
	_      = iota
	bronze = 10 * iota
	silver = 10 * iota
	gold   = 10 * iota
)

func main() {

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// 	second example

	fmt.Println(bronze)
	fmt.Println(silver)
	fmt.Println(gold)
}
