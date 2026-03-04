package excercise3trietree

import (
	"unicode"
)

func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	n := len(letters)
	for i := 0; i < n/2; i++ {
		if letters[i] != letters[n-1-i] {
			return false
		}
	}
	return true

}
