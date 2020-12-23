/***************************************************************
CHANNEL with SELECT
- SELECT
	-- While receiving some values form channel with certain
	condition or from certain channel select is used to determine
	which channel's value needs to be received.
***************************************************************/

package main

import (
	"fmt"
)

func main() {
	//create channels
	frontend := make(chan string)
	backend := make(chan string)
	quit := make(chan string)

	//send some values to channels with goroutine
	go send(frontend, backend, quit)

	//receive some values from channels
	receive(frontend, backend, quit)
}

func send(f, b, q chan<- string) {
	data := []string{"React", "NodeJS", "Vue", "FLask"}
	for i := 0; i < len(data); i++ {
		if i%2 == 0 {
			//send value to channel f
			f <- data[i]
		} else {
			b <- data[i]
		}
	}
	q <- "finished"
}

func receive(f, b, q <-chan string) {
	for {
		//using select to choose certain channel
		//that the value need to be received
		select {
		//if the value comes from channel called "f"
		//then execute the code
		case v := <-f:
			fmt.Println("frontend dev: ", v)
			//if the value comes from channel called "b"
			//then execute the code
		case v := <-b:
			fmt.Println("backend dev: ", v)
			//if the value comes from channel called "q"
			//then execute the code
		case v := <-q:
			fmt.Println("This program is ", v)
			return // finished the execution
		}
	}
}
