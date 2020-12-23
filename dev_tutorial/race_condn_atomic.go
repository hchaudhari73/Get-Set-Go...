/***************************************************************
USING ATOMIC to solve RACE CONDITION
***************************************************************/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64
var wg sync.WaitGroup

func main() {
	total := 10
	wg.Add(total)
	for i := 1; i <= total; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter", counter)
}
