/**************************************************************
ARRAY, SLICE and MAP
ARRAY
- Array isn't recommended to be used but it's good to know the
  characteristics of array in Go. Because using certain size,
  the size or length can't be changed automatically.

  Data stored must have a same type.

- KEYWORD
	-- data := [5]int{1, 2, 3, 4, 5}

	// empty array
	-- data := [2]int{}
	   data[0] = 1
	   data[1] = 2

SLICE
- Slice is a data structure that works like array but has a
  flexible size or length.

  Data stored must have a same type.

- SYNTAX
	-- data := []int{1, 2, 3, 4, 5} // size isn't needed.
	-- make([]int, 3, 6)

- INDEXING / SLICING
	-- data := [][]int{{}{}{}}
	-- data[1:3] -> data[1], data[2]

- APPEND
	-- data := []int{1, 2, 3}
	-- data = append(data, 4, 5, 6)
**************************************************************/

package main

import "fmt"

/*
func main() {
	data := [5]string{"NodeJS", "PHP", "Python", "GO"}

	for _, value := range data {
		fmt.Println(value)
	}
}
*/

func main() {
	data := make([]int, 3, 6) // empty array
	//the length of slice is 3
	//the capacity of slice is 6
	data[0] = 1
	data[1] = 2
	data[2] = 3

	fmt.Println("Length of slice =", len(data))
	fmt.Println("Capacity of slice =", cap(data))
	for _, value := range data {
		fmt.Println(value)
	}

	data = append(data, 4, 5, 6) //insert new items
	fmt.Println("Added items ", data)
	fmt.Println("Length of data after slice =", len(data))
	fmt.Println("Capacity of slice =", cap(data))
}
