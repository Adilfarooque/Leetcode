package main

func shortestToChar(s string, c byte) []int {
	res := make([]int, len(s))
	prev := -len(s)
	next := len(s) * 2
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			prev = i
		}
		res[i] = i - prev
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			next = i
		}
		res[i] = min(res[i], next-i)
	}
	return res
}
