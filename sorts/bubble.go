package main

import (
	"fmt"
	"math/rand"
)

func bubble_sort_ip(nums *[]int) {
	for i := 0; i < len(*nums); i++ {
		swap_took_place := false

		for j := 0; j < len(*nums)-1-i; j++ {
			if (*nums)[j] > (*nums)[j+1] {
				temp := (*nums)[j]
				(*nums)[j+1] = (*nums)[j]
				(*nums)[j] = temp

				swap_took_place = true
			}
		}

		if !swap_took_place {
			// end early if no swaps happened
			return
		}
	}

	// NOTE: both indices start at 0 and end at the loop end
	// NOTE: if we do - i in inner loop, we avoid unnecessary comparisons to values
	// 			 already in their correct spots
	//  		 i.e. AFTER ONE ITERATION, the LARGEST UNSORTED NUM MOVES TO THE END
	// NOTE: if we go through entire thing and no swaps happened, we can end premature
	// 		   we can implement this with a swap tracker variable in the inner loop
}

func main() {
	// [deprecated]
	// seed rng with current time in unix time
	// unix time is amt. of seconds elapsed since jan 1, 1970
	// rand.Seed(time.Now().UnixNano())

	// seed with my value for deterministic results
	// rand.New(rand.NewSource(67))
	// this creates diff results, each time, TODO: why???

	// populate SLICE with random nums
	// ...slices and arrays are different!!!
	nums := make([]int, 32)
	for i := 0; i < len(nums); i++ {
		nums[i] = rand.Intn(100) // random between 0 and 100 exclusive (so 0-99)
	}

	bubble_sort_ip(&nums)

	fmt.Printf("Sorted nums: ")
	for _, n := range nums {
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}

/*
	> o(n^2) complexity in worst case... yikes
	> this is a STABLE sort
*/
