package main

func isMonotonic(nums []int) bool {
	increment, decrement := true, true
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			increment = false
		}
		if nums[i] < nums[i+1] {
			decrement = false
		}
	}
	return increment || decrement

}
