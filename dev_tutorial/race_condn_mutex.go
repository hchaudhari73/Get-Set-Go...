/***************************************************************
USING MUTEX to solve RACE CONDITION
- MUTEX : Mutual Exclusion
	- Mechanism that allows only one goroutine to run the critical
	section (a code that has a potential of race condition) to
	prevent from race condition.
***************************************************************/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var counter int = 0
var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	var total int = 10
	wg.Add(total)
	for i := 0; i < total; i++ {
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
