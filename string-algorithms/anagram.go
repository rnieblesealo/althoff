package main

import (
	"fmt"
	"strings"
)

/*
	ANAGRAM DETECTION
	approach:
		sort the string!

	going to FIRST TRY WITH WHAT I KNOW then do it with suggested tools
*/

func sort_string(a string) string {
	/*
		string insertion sort
		move letter down until what's before it is <= to it
	*/

	// clean up
	a = strings.ReplaceAll(a, " ", "")
	// MIGHT NEED TO DO OTHER STUFF if processing weirder input

	// make into rune array for mutability
	// strings are immutable in go!
	s := []rune(a)
	for i := range s {
		if s[i] >= 'A' && s[i] <= 'Z' {
			s[i] += 32 // convert to lowercase
		}
	}

	// fmt.Printf("LOWER: %s\n", string(s))

	// do actual insertion sort
	for i := 1; i < len(s); i++ {
		val := s[i] // get ascii code of a (this will be a byte)

		for i > 0 && s[i-1] > val { // while what's before is greater than the current value, shift it down
			// NOTE: we can compare these without cast since they're bytes

			s[i] = s[i-1] // move the previous value up
			i -= 1        // continue down
		}

		s[i] = val // write our value at its correct location
	}

	// fmt.Printf("SORTED: %s\n", string(s))

	return string(s)
}

func is_anagram(a string, b string) bool {
	// sort both strings then compare
	return sort_string(a) == sort_string(b)
}

func main() {
	wa := "Emperor Octavian"
	wb := "Captain over Rome"

	fmt.Printf("%s : %s\n", wa, wb)

	if is_anagram(wa, wb) {
		fmt.Println("YES ANAGRAMS!")
	} else {
		fmt.Println("NOT anagrams")
	}
}
