/***************************************************************
CUSTOM ERROR HANDELING USING STRUCT
***************************************************************/
package main

import (
	"fmt"
	"log"
	"math"
)

type cust_err struct {
	a float64
}

//Implements Error() method so the cust_err struct is a type of
//error
//Basically cust_err struct is a error type

func (c *cust_err) Error() string {
	return fmt.Sprintf("cannot find sqrt of negative number %v", c.a)
}

func main() {
	value, err := sqrt(16)
	if err != nil {
		log.Println("Error occured: ", err)
	} else {
		fmt.Printf("sqrt %d = %f", 16, value)
	}
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, &cust_err{n}
	}
	return math.Sqrt(n), nil
}
