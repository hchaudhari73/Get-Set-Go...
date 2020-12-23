/***************************************************************
PIPELINE using FAN IN
- Fan In pattern is a concurrency pattern that takes soem inputs
and used it into one channel.
- Fan In pattern that takes some inputs and used it into one channel.
- It works like multiplexer.
***************************************************************/

package main

import ( 
	"fmt"
	"sync"
)

func main() {
	ints := generate()
	results := make(chan int)
	go average (ints, results)
	for  v:= range results {
		fmt.Println("Average:", v)
	}
}

//send some values to the channel
func generate() <- chan []int {
	//create a channel
	out := make(chan []int)
	go func() {
	//insert some values
	data := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		}
	for _, v := range data {
		out <- v
	}
	close(out)
	}()
	return out
}

func average(i <-chan []int, receiver chan int ) {
	//init waitgroup
	var wg sync.WaitGroup
	//add 1 goroutine
	wg.Add(1)
	
	//launch a goroutine
	go func() {
		//receive values from channel called i
		for  v := range i{
		receiver <- avg(v)
		}
	wg.Done()
	}()
	//wait until goroutine is finished
	wg.Wait()
	close(receiver)
}  

//function for calculating avg
func avg(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	results := total / len(nums)
	return results
}














