package main

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	BUCKET_COUNT = 16
)

/*
	CHALLENGE: research a sort that isn't in the book and write it!
		i choose BUCKET SORT because i have seen it in some problems and it seems useful
*/

/*
	BUCKET SORT:
		categorize elements of array into buckets
		THEN sort each bucket individually with a different sort or w recursive bucketsort
*/

func bucketsort_ip(nums *[]int, k int) { // k is the AMOUNT OF BUCETS TO USE
	// allocate buckets

	buckets := make([]([]int), k) // array of slices (i.e. VECTORS)

	// get maximum key value

	maxkey := -1
	for _, n := range *nums {
		if n > maxkey {
			maxkey = n
		}
	}

	maxkey = 1 + maxkey // NOTE: must add 1 to avoid offbounds bs! DO NOT FORGET

	// put shit into buckets

	for i := 0; i < len(*nums); i++ {
		bucket_loc := int(math.Floor(float64(k * (*nums)[i] / maxkey))) // all casts here avoid errors
		fmt.Printf("BUCKET LOC: %d\n", bucket_loc)

		buckets[bucket_loc] = append(buckets[bucket_loc], (*nums)[i])
	}

	// sort stuff in each bucket individually

	for _, bucket := range buckets {
		nextsort(&bucket)

		// DEBUG: print the bucket items...
		for _, value := range bucket {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}

	// concatenate the buckets
	// "i" should finish at the original array length

	i := 0
	for _, bucket := range buckets {
		for _, value := range bucket {
			(*nums)[i] = value
			i++
		}
	}

	// done!
}

func nextsort(nums *[]int) { // the individual sort for smaller buckets; using insertion
	// move down until in correct slot
	for i := 1; i < len(*nums); i++ {
		val := (*nums)[i]
		for i > 0 && (*nums)[i-1] > val {
			(*nums)[i] = (*nums)[i-1] // move down
			i = i - 1
		}
		(*nums)[i] = val // write val to correct spot
	}
}

func main() {
	nums := make([]int, 64)
	for i := 0; i < len(nums); i++ {
		nums[i] = rand.Intn(128) // between 0-127
	}

	bucketsort_ip(&nums, BUCKET_COUNT) // use

	fmt.Println("BUCKET SORTED array: ")
	for _, n := range nums {
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}

/*
	HOLY SHIT

	i got this to work with the pseudocode ONLY first try
		im him
		SOURCE: https://en.wikipedia.org/wiki/Bucket_sort

	RUNTIME ANALYSIS

	WORST CASE:
		everything goes in ONE BUCKET in which case the nextsort algo determines speed
			here since i used ins. sort it would be o(n^2)

	AVERAGE CASE:
		O(n) to find max key
		O(n) to categorize everything

		...so far O(n + n) = O(2n) = O(n)

		if we use INSERTION SORT for each bucket O(n^2),

		it's O(n + n^2)
			(wikipedia says inssort accounts for n^2/k, i assume in average, so let's use that...)

		...so far O(n + n^2/k)

		the CONCATENATION STEP requires O(k) time where k = bucket count
			(is this real in the way im doing it???)

		TOTAL:

			O(n + n^2/k + k)
				MIDDLE TERM corresponds to the nextsort algo!
				K IS THE AMOUNT OF BUCKETS WE DECIDED TO USE
*/
