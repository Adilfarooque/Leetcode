package main

func largeGroupPositions(s string) [][]int {
	var res [][]int
	var i, j int
	for i < len(s) {
		j = i
		for j < len(s) && s[j] == s[i] {
			j++
		}
		if j-i >= 3 {
			res = append(res, []int{i, j - 1})
		}
		i = j
	}
	return res
}
