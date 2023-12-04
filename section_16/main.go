package main

import (
	"fmt"
)

func main() {
	last_first := map[string][]string{
		"Bond_James":       {"shaken, not stirred", "martinis", "fast cars"},
		"Moneypenny_Jenny": {"intelegence", "literature", "computer science"},
		"no_dr":            {"cats", "ice cream", "sunsets"},
	}

	for k, v := range last_first {
		fmt.Printf("These are map keys: %v\n", k)
		fmt.Printf("These are map values: %v\n", v)
		for index, value := range last_first[k] {
			fmt.Printf("For key %v, the value at index %v is %v\n",
				k, index, value)
		}
	}
}
