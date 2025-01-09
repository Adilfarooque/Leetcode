package main

import "strings"

//2185. Counting Words With a Given Prefix

func prefixCount(words []string, pref string) int {
	count := 0
	for _, word := range words {
		if strings.HasPrefix(word, pref) {
			count++
		}
	}
	return count
}

