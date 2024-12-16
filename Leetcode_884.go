package main

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	wordsCount := make(map[string]int)
	word1 := strings.Split(s1, " ")
	word2 := strings.Split(s2, " ")
	for _, word := range word1 {
		wordsCount[word]++
	}
	for _, word := range word2 {
		wordsCount[word]++
	}
	res := []string{}
	for word, count := range wordsCount {
		if count == 1 {
			res = append(res, word)
		}
	}
	return res
}
