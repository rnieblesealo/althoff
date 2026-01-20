package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// alright we rollin this one on main
	// cuz im gungry

	s := "pass the 2oobie to the left hand 5ide" // 5 is rightmost

	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToLower(s)

	var _s []byte
	for _, letter := range s {
		if unicode.IsDigit(letter) {
			_s = append(_s, byte(letter))
		}
	}

	if len(_s) > 0 {
		fmt.Printf("LAST DIGIT: %s\n", string(_s[len(_s)-1])) // print the last
	}
}
