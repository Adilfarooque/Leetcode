package main

//serach in rotated sorted array
func search(nums []int, target int) bool {
	for _, value := range nums {
		if target == value {
			return true
		}
	}
	return false
}
