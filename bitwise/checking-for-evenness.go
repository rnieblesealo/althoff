package main

import (
	"fmt"
)

func is_even(n int) bool {
	// since the number 1 is the LSB in binary (rightmost value)
	// and all other place values are a power of 2,
	// only this 1 can throw the thing off (powers of 2 are all otherwise also divisible by 2)

	// therefore if we just check that the lsb is set, then we know the thing is even

	return ^n&1 == 1
}

func main() {
	n := []int{1, 2, 3, 4, 5, 6}
	for _, n := range n {
		if is_even(n) {
			fmt.Printf("%d is even\n", n)
		} else {
			fmt.Printf("%d is not even\n", n)
		}
	}
}

// THIS IS REALLY COOL
