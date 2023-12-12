package main

import (
	"fmt"
	"os"
)

func main() {
	result, err := Add(3, 1)
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("The result is %v", result)
}

// Add takes two integers and returns the sum of them and a nil error if
// everything goes right. The integers have to be non-negative, otherwise Add
// will return a zero and an error!
func Add(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, fmt.Errorf("a and b need to be non-negative integers")
	}

	return a - b, nil
}
