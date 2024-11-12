package main

func repeatedNTimes(nums []int) int {
	exists := make(map[int]bool)
	for _, value := range nums {
		if exists[value] {
			return value
		}
		exists[value] = true
	}

	return -1

}
