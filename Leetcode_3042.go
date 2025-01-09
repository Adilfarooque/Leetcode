package main

import "strings"

// 3042. Count Prefix and Suffix Pairs I
func countPrefixSuffixPairs(words []string) int{
    count := 0
    for i, word := range words {
        for j, w := range words {
            if i != j && strings.HasSuffix(word, w[:len(word)-1]) {
                count++
            }
        }
    }
    return count
}
