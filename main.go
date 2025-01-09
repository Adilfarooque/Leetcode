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
	fmt.Println("Square of sorted array:", SortedSquares(nums))

	//Largest Perimeter
	pNumbers := []int{2, 1, 2, 4, 3}
	fmt.Println("Largest perimeters :", largestPerimeter(pNumbers))

	//Repeated N Times
	NTimes := []int{2, 1, 2, 5, 3, 2}
	fmt.Println("Repeated elements :", repeatedNTimes(NTimes))

	//Valid mountain Array
	mountain := []int{0, 3, 2, 1}
	fmt.Println("Valid mountain :", validMountainArray(mountain))

	//Min Deletion Size
	strs := []string{"cba", "daf", "ghi"}
	fmt.Println("Min Deletion Size :", MinDeletionSize(strs))

	//dIStringMatch
	s := "III"
	fmt.Println("DIStringMatch :", diStringMatch(s))
	s = "IDID"
	fmt.Println("DIStringMatch :", diStringMatch(s))

	//Reverse only letters
	s = "a-bC-dEf-ghIj"
	fmt.Println("Reverse only letters :", reverseOnlyLetters(s))
	s = "Test1ng-Leet=code-Q!"
	fmt.Println("Reverse only letters :", reverseOnlyLetters(s))

	//Sort Array by Parity
	nums = []int{3, 1, 2, 4}
	fmt.Println("Sort Array by Parity :", SortArrayByParity(nums))

	//Is Monotonic
	nums = []int{1, 2, 2, 3}
	fmt.Println("Is Monotonic :", isMonotonic(nums))
	nums = []int{6, 5, 4, 4}
	fmt.Println("Is Monotonic :", isMonotonic(nums))

	//uncommonFromSentences
	s1 := "this apple is sweet"
	s2 := "this apple is sour"
	fmt.Println("uncommonFromSentences :", uncommonFromSentences(s1, s2))

	//buddyStrings
	test := "ab"
	goal := "ba"
	fmt.Println("buddyStrings :", buddyStrings(test, goal))
	test = "ab"
	goal = "ab"
	fmt.Println("buddyStrings :", buddyStrings(test, goal))
	s = "aa"

	//backspaceCompare
	test = "ab#c"
	goal = "ad#c"
	fmt.Println("backspaceCompare :", backspaceCompare(test, goal))
	test = "ab"
	goal = "c#d#"
	fmt.Println("backspaceCompare :", backspaceCompare(test, goal))
	test = "a#c"
	goal = "b"
	fmt.Println("backspaceCompare :", backspaceCompare(test, goal))

	//largegroupPositions
	s = "abbxxxxzzy"
	fmt.Println("largegroupPositions :", largeGroupPositions(s))
	s = "abc"
	fmt.Println("largegroupPositions :", largeGroupPositions(s))

	//shortestToChar
	s = "loveleetcode"
	c := byte('e')
	fmt.Println("shortestToChar :", shortestToChar(s, c))

	//addedInteger
	nums1 := []int{2, 6, 4}
	nums2 := []int{9, 7, 5}
	fmt.Println("addedInteger :", addedInteger(nums1, nums2))

	//search
	nums = []int{2, 5, 6, 0, 0, 1, 2}
	target := 0
	fmt.Println("search :", search(nums, target))
	nums = []int{2, 5, 6, 0, 0, 1, 2}
	target = 3
	fmt.Println("search :", search(nums, target))

	//maxScore
	s = "011101"
	fmt.Println("maxScore :", maxScore(s))
	s = "00111"
	fmt.Println("maxScore :", maxScore(s))
	s = "1111"
	fmt.Println("maxScore :", maxScore(s))

	//2559. Count Vowel Strings in Ranges
	words := []string{"aba", "bcb", "ece", "aa", "e"}
	queries := [][]int{{0, 2}, {1, 4}, {1, 1}}
	fmt.Println("Vowel Strings :", vowelStrings(words, queries))

	//2270. Number of Ways to Split Array
	nums = []int{10, 4, -8, 7}
	fmt.Println("Number of Ways to Split Array :", waysToSplitArray(nums))
	nums = []int{2, 3, 1, 0}
	fmt.Println("Number of Ways to Split Array :", waysToSplitArray(nums))

	//prefixCount
	words = []string{"pay", "attention", "practice", "attend"}
	pref := "at"
	fmt.Println("prefixCount :", prefixCount(words, pref))
	
}
