package main

import (
	"fmt"
	"strings"
)

func search(sent string, word string) bool {
	return strings.Contains(sent, word)
}

func main() {
	sentence := "I am thinking to start a startup."
	word := "start"
	if search(sentence, word) {
		fmt.Printf("Contains %s\n", word)
	} else {
		fmt.Printf("Does not contains %s\n", word)
	}
}
