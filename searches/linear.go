package main

// FIXME: Get rid of no packages found for open file warning...

import (
	"fmt"
)

func linear_search(nums []int, target int) bool {
	for _, n := range nums {
		if n == target {
			return true
		}
	}

	return false
}

func main() {
	nums := []int{12, 3, 67, 67, 67, 3, 4, 12, 1}

	fmt.Printf("sixseven? ")

	res := linear_search(nums, 67) // NOTE: Needed to make this a slice... What's up with that?
	if res == true {
		fmt.Printf("YES\n")
	} else {
		fmt.Printf("NO :(\n")
	}
}

/*
	some notes on golang:

	> no ternary, if else is idiomatic
	> main doesn't return anything

	on LINEAR SEARCH:

	> o(n/2) average complexity
	> o(n) worst case complexity
	> o(1) best case

	> use when order can't be leveraged (this algo dgaf about order)
*/
