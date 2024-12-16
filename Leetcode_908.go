package main

func SmallestRangeI(nums []int, k int) int {
	min, max := nums[0], nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	if max-min <= 2*k {
		return 0
	}
	return max - min - 2*k
}
