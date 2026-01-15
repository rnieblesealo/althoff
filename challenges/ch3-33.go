package main

import (
	"fmt"
	"sort"
)

/*
challenge: given a list of words, write a func that binary searches for a word in that list
	i will assume each word will begin with a distinct letter and is in lowercase
	otherwise the challenge becomes more complicated which i think is not intended
		(but would be very fun to tackle!!!)
*/

func charbin_search(words []string, target string, l int, r int) bool {
	if l > r {
		return false
	}

	ord_target := int(target[0])

	mid := int(l + (r-l)/2)
	if words[mid] == target {
		return true
	} else {
		ord_mid := int(words[mid][0])
		if ord_mid < ord_target {
			return charbin_search(words, target, mid+1, r)
		} else {
			return charbin_search(words, target, l, mid-1)
		}
	}
}

func main() {
	words := []string{"banana", "apple", "orange", "grape", "kiwi"}
	sort.Strings(words)

	fmt.Printf("banana? ")

	res := charbin_search(words, "banana", 0, len(words)-1)
	if res == true {
		fmt.Println("YES")
	} else {
		fmt.Println("NO :/")
	}
}
