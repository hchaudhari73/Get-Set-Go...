package main

import (
	"fmt"
)

func avg(marks ...int) (int, int) {
	total := 0
	for _, value := range marks {
		total += value
	}
	average := total / len(marks)
	return total, average
}

func main() {
	total, average := avg(5, 5, 5)
	fmt.Println(total, average)
}
