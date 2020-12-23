/***************************************************************
CHANNEL
- Create a concurrent program in Go has a unique approch.
- The approch is share memory by communicating instead communicate
by sharing memory (which is not allowed in Go).
- This approch can be done using channel that can be used by
goroutines to communicate with each other.
***************************************************************/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	//attach or send a value into channel
	//put a value into channel
	//wrapped inside goroutine
	go func() {
		c <- 43
	}()

	//retrieve a value from channel
	fmt.Println("Value of channel c: ", <-c)
}
