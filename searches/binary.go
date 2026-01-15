package main

import (
	"fmt"
	"sort"
)

func binary_search(nums []int, target int, l int, r int) bool {
	if l > r {
		return false
	}

	mid := int((l + r) / 2) // overflow-safe way: int(l + (r - l) / 2)
	if nums[mid] == target {
		return true
	} else if nums[mid] < target {
		return binary_search(nums, target, mid+1, r)
	} else {
		return binary_search(nums, target, l, mid-1)
	}
}

func main() {
	nums := []int{12, 3, 67, 67, 67, 3, 4, 12, 1}
	sort.Ints(nums)

	fmt.Printf("sixseven? ")

	res := binary_search(nums, 67, 0, len(nums)-1)
	if res == true {
		fmt.Printf("YES!\n")
	} else {
		fmt.Printf("NOOO")
	}
}

/*
	notes on BINARY SEARCH:

	> o(n) best case (find target right away)
	> o(logn) average AND worst case
		we always HALVE our load!

	> use when sorting CAN be leveraged
		...if you painfully sort data once, it's VERY SLOW for this one time but the savings can be huge!
*/
