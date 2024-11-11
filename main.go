package main

import (
	"fmt"
)

func main() {
	//Move All zeros
	numbers := []int{1, 2, 0, 1, 3, 0, 22, 0, 4, 6}
	MoveAllzero(numbers)
	fmt.Print(numbers)

	//Reverse String
	str := "Hello, World!"
	revrese := ReverseString(str)
	fmt.Println(revrese)

	byteStr := []byte{'h', 'e', 'l', 'l', 'o'}
	ReverseBytesOfString(byteStr)
	fmt.Println(string(byteStr))

	//Sorted Squares
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Print("Square of sorted array:",SortedSquares(nums))
}
