package main


//2270. Number of Ways to Split Array

func waysToSplitArray(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	count := 0
	left := 0
	for i := 0; i < len(nums)-1; i++ {
		left += nums[i]
		if left >= sum-left {
			count++
		}
	}
	return count
}