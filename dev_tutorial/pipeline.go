/***************************************************************
PIPELINE
- Pipeline pattern is basically works like a pipe that connect
each other. 
- MAJOR STEPS
	-- Get values from channels
	-- Perform some operations with the value
	-- Send values to channel so the value can be consumed
	or received.
***************************************************************/

package main

import (
	"fmt"
)

func main() {
	ints := generate()
	results := average(ints)
	for v := range results {
		fmt.Println("Average:", v)
	}
}

//STEP 1: Get the values from channel
func generate() <-chan []int {
	out := make(chan []int)
	go func() {
		//insert some data into channel
		data := [][]int{{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9}}
		for _, v := range data {
			out <- v
		} 
		//close the channel which means the operation
		//to channel is finished
		close(out)
	}()
	return out
}	

//STEP 2: Perform operation with the value, in this case is avg caln.
func average(i <-chan []int) <-chan int {
	//create a channel
	out := make(chan int)
	go func() {
		//receive value from a channel
		//the receive value is []int
		for nums := range i {
			//then cal value's average
			//step 3: Send values to channel
			out <- avg(nums)
		}
		close(out)
	}()
	return out
} 

//function for calculating average of numbers
func avg(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	result := sum / len(nums)
	return result
}
















