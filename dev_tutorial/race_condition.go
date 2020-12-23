/***************************************************************
RACE CONDITION
- Condition when write and read operation occur in the same
variable at the same time.
- Makes the value of certain variable inconsistent and intrduces
bug in a code.
- SYNTAX
	-- go run -- race
- Alternate Defination:
	A race condition is an undersirable situation that occurs
when a device or system attempts to perform two or more operations
at the same time, but becaus of the nature of the device or system,
the operations must be done in the proper sequence to be done
correctly.
***************************************************************/

package main

import (
	"fmt"
)

var counter int = 0

func main() {
	total := 10
	for i := 0; i < total; i++ {
		go func() {
			v := counter
			// runtime.Gosched() //can be replaced with time.sleep()
			v++
			counter = v
		}()
		fmt.Println("Counter: ", counter)
	}
}
