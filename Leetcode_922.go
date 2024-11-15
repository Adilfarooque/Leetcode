package main

func SortArrayByPartyII(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	even, odd := 0, 1
	for _, value := range nums {
		if value%2 == 0 {
			res[even] = value
			even += 2
		} else {
			res[odd] = value
			odd += 2
		}
	}
	return res
}
