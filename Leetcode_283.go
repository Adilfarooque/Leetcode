package main

import "fmt"

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
func main() {
	numbers := []int{1, 2, 0, 1, 3, 0, 22, 0, 4, 6}
	MoveAllzero(numbers)
	fmt.Println(numbers)
}
