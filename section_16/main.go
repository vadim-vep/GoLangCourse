package main

import (
	"fmt"
)

// declaring map and adding keys and values
// in this case - keys are type string and values are slices with strings inside
func main() {
	last_first := make(map[string][]string)
	last_first["Bond_James"] = []string{"shaken, not stirred", "martinis", "fast cars"}
	last_first["Moneypenny_Jenny"] = []string{"intelegence", "literature", "computer science"}
	last_first["no_dr"] = []string{"cats", "ice cream", "sunsets"}

	for k, v := range last_first {
		fmt.Printf("These are map keys: %v\n", k)
		fmt.Printf("These are map values: %v\n", v)
		for index, value := range v {
			fmt.Printf("For key %v, the value at index %v is %v\n",
				k, index, value)
		}
	}

	//adding another value to the map
	last_first["Ian Flemming"] = []string{"steaks", "cigars", "espionage"}
	fmt.Println(last_first)

	//deleting from the map
	delete(last_first, "no_dr")
	for k, v := range last_first {
		fmt.Printf("The key after deletion is %v\n", k)
		fmt.Printf("The value after deletion is %v\n", v)
	}

}
