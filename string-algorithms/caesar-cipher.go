package main

import (
	"fmt"
	"strings"
)

// k = amount of spaces to shift
// key is to use MOD to wrap around
// mod 26 since there are 26 letters
// transform the value into the ascii range
// for simplicity, we'll only take lowercase from a-z

func encrypt(s string, k int) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")

	s_encrypted := make([]byte, len(s))

	for i := range len(s) {
		c := s[i] - 'a'
		c = (c + byte(k)) % 26

		s_encrypted[i] = byte(c + 'a')
	}

	return string(s_encrypted)
}

func decrypt(s string, k int) string {
	// WARNING: this assumes that the string is PROPERLY FORMATTED! (no spaces, all lower)

	s_decrypted := make([]byte, len(s))

	for i := range len(s) {
		c := int(s[i] - 'a')
		c = c - k

		if c < 0 {
			c = (26 + c) // <---- CAREFUL HERE ON WRAP! since c is negative we need to ADD IT
		}

		s_decrypted[i] = byte(c + 'a')
	}

	return string(s_decrypted)
}

func main() {
	s := "hello my name is caesar"

	s_enc := encrypt(s, 3)
	s_dec := decrypt(s_enc, 3)

	fmt.Printf("%s -e--> %s -d-->%s\n", s, s_enc, s_dec)
}
