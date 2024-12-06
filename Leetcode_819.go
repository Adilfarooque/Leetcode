package main

import (
	"strings"
	"unicode"
)

func mostCommonWords(paragraph string, banned []string) string {
	ban := map[string]bool{}
	for _, s := range banned {
		ban[s] = true
	}
	words := strings.FieldsFunc(paragraph, func(r rune) bool {
		return !unicode.IsLetter(r)
	})
	count := map[string]int{}
	for _, w := range words {
		w = strings.ToLower(w)
		if !ban[w] {
			count[w]++
		}
	}
	maxCount := 0
	ans := ""
	for w, c := range count {
		if c > maxCount {
			maxCount = c
			ans = w
		}
	}
	return ans
}
