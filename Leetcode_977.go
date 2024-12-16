package main

import "sort"

//Square of a sorted array
func SortedSquares(nums []int) []int {
	result := []int{}
	for _, value := range nums {
		result = append(result, value*value)
	}
	sort.Ints(result)
	return result
}
