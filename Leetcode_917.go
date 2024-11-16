package main

import (
	"unicode"
)

func reverseOnlyLetters(s string) string {
	left, right := 0, len(s)-1
	str := []byte(s)
	for left < right {
		if !unicode.IsLetter(rune(s[left])) {
			left++
			continue
		}
		if !unicode.IsLetter(rune(s[right])) {
			right--
			continue
		}
		str[left], str[right] = str[right], str[left]
		left++
		right--
	}
	return string(str)
}
