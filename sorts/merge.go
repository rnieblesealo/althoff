package main

import (
	"fmt"
	"math/rand"
)

func mergesort_ip(nums *[]int) {
	if len(*nums) <= 1 {
		return
	}

	// break apart

	mid := int(len(*nums) / 2)

	lhalf := (*nums)[0:mid]
	rhalf := (*nums)[mid:len(*nums)]

	// recur on these halves...
	// WARNING: this might break and crash and burn
	// UPDATE: it didn't!
	// we're only overwriting the big array once our splits are done, makes sense
	mergesort_ip(&lhalf)
	mergesort_ip(&rhalf)

	l := 0 // where we are on LEFT subarray
	r := 0 // where we are on RIGHT subarray
	k := 0 // where we are on MERGED array

	// go over subarrays
	for l < len(lhalf) && r < len(rhalf) {
		// compare and merge in SORTED order
		// advance ONLY the index of the subarray we copied from!
		if lhalf[l] <= rhalf[r] {
			(*nums)[k] = lhalf[l]
			l++
		} else {
			(*nums)[k] = rhalf[r]
			r++
		}

		// advance k ALWAYS since we always copy something to it
		k++
	}

	// if an array is depleted, PASTE THE REMAINING STUFF from the other one

	for l < len(lhalf) {
		(*nums)[k] = lhalf[l]
		l++
		k++
	}

	for r < len(rhalf) {
		(*nums)[k] = rhalf[r]
		r++
		k++
	}
}

func main() {
	// make random number array
	nums := make([]int, 64)
	for i := 0; i < len(nums); i++ {
		nums[i] = rand.Intn(127)
	}

	// mergesort in place
	mergesort_ip(&nums)

	// print it
	fmt.Println("Mergesorted array: ")
	for _, n := range nums {
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}

/*
	break in half until individual
	then merge in sorted order
*/
