package main

//Move All zero into the end
func MoveAllzero(nums []int) {
	fast, slow := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[fast], nums[slow] = nums[slow], nums[fast]
			slow++
		}
		fast++
	}
}
