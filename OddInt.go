package main

import (
	"fmt"
)

func main() {
	input := []int{3, 4, 4, 3, 3}

	fmt.Println(input)
	allKeys := make(map[int]int)

	for i := 0; i < len(input); i++ {

		value, _ := allKeys[input[i]]
		value++
		allKeys[input[i]] = value
	}

	for key := range allKeys {
		value, _ := allKeys[key]
		if value%2 != 0 {
			fmt.Println(key, "is Odd number")
		}
	}
}
