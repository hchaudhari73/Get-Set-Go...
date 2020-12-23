/**************************************************************
INTERATION WITH FOR RANGE
- KEYWORD
	-- range data

- FOR
	-- for index, value := range data {....}

- INDEXING OF ARRAY
	-- v[i]
**************************************************************/

package main

import (
	"fmt"
)

/*
func main() {
	data := [5]int{1, 2, 3, 4, 5} //declares an array of int
	for i, v := range data {
		fmt.Println("index ", i, "value ", v)
	}
}
*/

/*
func main() {
	data := [5]int{1, 2, 3, 4, 5}
	total := 0
	for _, v := range data {
		total += v
	}
	fmt.Println(total)
}
*/

func main() {
	data := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	for _, v := range data {
		fmt.Println("Data: ", v)
		for i := 0; i < len(v); i++ {
			fmt.Println(v[i])
		}
	}
}
