package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 42
	data := float64(a)
	b := "42"
	from_string, _ := strconv.Atoi(b)
	to_string := strconv.Itoa(a)

	fmt.Printf("%d, %T\n", a, a)
	fmt.Printf("%s %T\n", b, b)
	fmt.Printf("%f %T\n", data, data)
	fmt.Printf("%d %T\n", from_string, from_string)
	fmt.Printf("%s %T\n", to_string, to_string)
}
