/**************************************************************
ERROR HANDLING
- Error can be checked with
	-- fmt.Println("Error occured: ", err)
	-- log.Println(err)
	-- log.Fatalln(err) //shows date and time when the error occured
	-- panic(err)
**************************************************************/

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//reading a .txt file that doesn't exist
	_, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Error occured: ", err)
		log.Fatalln("Error occured: ", err)
	}
}
