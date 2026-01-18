package main

import (
	"fmt"
	"math/rand"
)

/*
	this one is kinda confusing but
	you just SHIFT the current value down until it's in the correct spot
		i.e. whatever's before it is less than it and what's after it is larger

	really easy
*/

func inssortip(nums *[]int) {
	for i := 1; i < len(*nums); i++ { // make sure to start at 1
		// get the current val
		temp := (*nums)[i]

		// move it down until before is LESS and after is MORE
		// ensure we don't go off bounds!
		for i > 0 && (*nums)[i-1] > temp {
			(*nums)[i] = (*nums)[i-1] // place prev to where we are now
			i = i - 1                 // decrement
		}

		// write the value to where it belongs
		(*nums)[i] = temp
	}
}

func main() {
	// random array
	nums := make([]int, 32)
	for i := 0; i < len(nums); i++ {
		nums[i] = rand.Intn(128) // between 0-127
	}

	// sort in place using ins!
	inssortip(&nums)

	fmt.Println("Insertion-sorted array: ")
	for _, n := range nums {
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}

/*
though this is o(n^2)
	(we might need to push down the entire list for every value),
for NEARLY SORTED DATA, it's fast!
	insertion sort is used in some high-performance algorithms
	(i think?) it's also great because it involves less cache misses? something like that

this sort is STABLE
*/
