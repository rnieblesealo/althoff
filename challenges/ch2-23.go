package main

import (
	"fmt"
)

// challenge: print the numbers 1 to 10 recursively

func _1to10(n int) {
	if n == 0 {
		return
	}

	_1to10(n - 1)

	fmt.Printf("%d\n", n)
}

func main() {
	_1to10(10)
}
