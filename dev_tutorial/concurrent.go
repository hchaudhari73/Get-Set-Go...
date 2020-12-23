/***************************************************************
CONCURRENCY
- Concurrency != Parallelism
- Concurrency is a mechanism when two or more taks can start,
run in overlapped time.
- Parallelism is a mechanism when many tasks can run at the
same time.
- Two main components when create a concurrent program in Go.
	-- goroutine
		-- light weight thread, created using go keyword
	-- channel

WAITGROUP
- Waitgroup is a mechanism that can be used to sync the go code.
- SYNTAX
	-- Add() : defines the number of goroutines that involved.
	-- Wait() : defines the wait condition in certain goroutine.
	-- Done() : defines the finish condition in certain goroutine.
	This means that the operation inside goroutine is finished.
***************************************************************/

package main

/*
func main() {
	fmt.Println("This prints out")
	//create a goroutine
}

func odd() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}
}
/*
OUTPUT -> This prints out
*/

/*
//initiate the waitgroup
var wg sync.WaitGroup

func main() {
	wg.Add(1)
	fmt.Println("This prints out")
	go odd()  //goroutine
	wg.Wait() //wait until odd() function is finished
}

func odd() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println(i)
	}
}
OUTPUT -> This prints out
1
3
5
7
9
*/
