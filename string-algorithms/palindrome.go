package main

import (
	"fmt"
	"strings"
)

func reverse_string(a string) string {
	/* copy string chars to rune array in reverse order */

	b := make([]byte, len(a))
	for i := range len(a) {
		b[len(a)-1-i] = a[i]
	}

	return string(b)
}

func is_palindrome(a string) bool {
	// clean up
	a = strings.ToLower(a)
	a = strings.ReplaceAll(a, " ", "")

	return a == reverse_string(a)
}

func main() {
	a := "sabas"

	fmt.Printf("'%s' is palindrome?\n", a)

	if is_palindrome(a) {
		fmt.Println("YES")
	} else {
		fmt.Println("no")
	}
}
