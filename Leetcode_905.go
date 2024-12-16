package main

func SortArrayByParity(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]%2 == 1 && nums[right]%2 == 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
		if nums[left]%2 == 0 {
			left++
		}

		if nums[right]%2 == 1 {
			right--
		}
	}
	return nums
}
