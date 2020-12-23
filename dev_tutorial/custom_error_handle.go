/**************************************************************
CUSTOM ERROR HANDLE
**************************************************************/
package main

import (
	"errors"
	"log"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	_, err := divide(4, 0)
	if err != nil {
		defer log.Fatalln("Error occured: ", err)
		log.Println("Error occured: ", err)
	}
}
