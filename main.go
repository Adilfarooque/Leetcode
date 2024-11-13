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
	fmt.Println("Square of sorted array:",SortedSquares(nums))

	//Largest Perimeter
	pNumbers := []int{2,1,2,4,3}
	fmt.Println("Largest perimeters :",largestPerimeter(pNumbers))


	//Repeated N Times
	NTimes := []int{2,1,2,5,3,2}
	fmt.Println("Repeated elements :",repeatedNTimes(NTimes))

	//Valid mountain Array
	mountain := []int{0,3,2,1}
	fmt.Println("Valid mountain :",validMountainArray(mountain))
}

