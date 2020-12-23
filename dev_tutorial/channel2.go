/***************************************************************
CHANNEL OPEN and CLOSE
***************************************************************/

package main

import (
	"fmt"
)

func main() {
	//create a channel with capacity of 10 ints
	c := make(chan int, 10)

	//put a value into channel
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		//close the channel
		//this means that the sending operation is finished
		close(c)
	}()
	//retrieve a value from channel using for loop
	for v := range c {
		fmt.Printf("Value of channel %d\n", v)
	}
	fmt.Println("Bye")
}
