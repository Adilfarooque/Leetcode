package main

//2559. Count Vowel Strings in Ranges

func vowelStrings(words []string, queries [][]int) []int {
	n := len(words)
	prefix := make([]int, n+1)
	for i, word := range words {
		prefix[i+1] = prefix[i]
		if isVowelString(word) {
			prefix[i+1]++
		}
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = prefix[q[1]+1] - prefix[q[0]]
	}
	return ans
}

func isVowelString(word string) bool {
	n := len(word)
	return isVowel(word[0]) && isVowel(word[n-1])
}

func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}
