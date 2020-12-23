/**************************************************************
MAP
- Key value pairs mechanism.
- Unique keys.package dev_tutorial
- KEYWORD
	-- data := map[string]string{
		"BMW" : "car",
		"vs code" : "IDE",
		"Python" : "Language"
	}
	-- scores := map[string][]int{
    	"Firstname":{9,8,9},
    	"Lastname":{7,6,5},
    	"Surname":{8,8,8}
	}
**************************************************************/
package main

import (
	"fmt"
)

func main() {
	map1 := map[string]string{
		"BMW":     "car",
		"vs code": "IDE",
		"Python":  "Language",
	}

	fmt.Println(map1["BMW"])
	fmt.Println(map1["vs code"])
	fmt.Println(map1["Python"])

	map2 := map[string]int{
		"BMW":     125,
		"vs code": 215,
		"Python":  654,
	}

	for key, value := range map2 {
		fmt.Printf("key %s \tvalue %d\n", key, value)
	}

	map3 := map[string][]int{
		"MI":  {100, 123, 201},
		"RBC": {120, 101, 189},
		"CSK": {210, 101, 98},
	}

	for key, value := range map3 {
		fmt.Println(key, value)
	}
}
