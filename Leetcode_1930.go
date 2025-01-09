package main

//1930. Unique Length-3 Palindromic Subsequences

func CountPalindromicSubsequence(s string) int {
	count := 0
	for i := 0; i < 26; i++ {
		first, last := -1, -1
		for j := 0; j < len(s); j++ {
			if int(s[j]-'a') == i {
				if first == -1 {
					first = j
				}
				last = j
			}
		}
		if first != -1 && last != -1 && first != last {
			seen := make(map[byte]bool)
			for j := first + 1; j < last; j++ {
				if !seen[s[j]] {
					seen[s[j]] = true
					count++
				}
			}
		}
	}
	return count
}

// func countPalindromicSubsequence(s string) int {
//     res, inf := 0, 1 << 31
//     first, last := make([]int, 26),  make([]int, 26) // 记录每个字符第一次出现的位置和最后一次出现的位置
//     for i := range first{
//         first[i] = inf
//     }
//     for i := range s {
//         index := int(s[i]) - int('a')
//         first[index], last[index] = min(first[index], i), i
//     }
//     for i := 0; i < 26;i++ {
//         if first[i] < last[i] {
//             substring := s[first[i] + 1 : last[i]]
//             set := make(map[rune]bool)
//             for _, v := range substring {
//                 set[v] = true
//             }
//             res += len(set)
//         }
//     }
//     return res
// }

func countPalindromicSubsequence(s string) int {
	// 思路：因为题目只要求长度为3，那么就意味着，该回文肯定会有一个中心字符，核心思想是只要在该字符两侧找到一个相同字符即可，
	// 维护两个数组pre[i]和suff[i],表示前缀的字符比特位(不包括当前位)，如果对应字符的比特位为1，表示前缀拥有该字符；后缀也如此
	// 那么以该字符为中心的回文子序列为：pre[i]&suff[i] 中为1的个数
	// 可能会出现重复，因此还需要维护一个[26]int arr的数组 ，arr[i]表示以第i个字符为中心的回文子序列个数(1的个数就是回文子序列的个数)
	pre, suff := make([]int, len(s)), make([]int, len(s))
	for i := 1; i < len(s); i++ {
		pre[i] = pre[i-1] | (1 << (s[i-1] - 'a'))
	}
	for i := len(s) - 2; i >= 0; i-- {
		suff[i] = suff[i+1] | (1 << (s[i+1] - 'a'))
	}
	arr := [26]int{} // arr表示以第i个字符为中心的回文子序列的情况
	for i := 1; i < len(s)-1; i++ {
		arr[s[i]-'a'] |= (pre[i] & suff[i])
	}
	bitCount := func(bit int) int {
		res := 0
		for i := 0; i < 32; i++ {
			if (bit >> i & 1) == 1 {
				res++
			}
		}
		return res
	}
	res := 0
	for i := 0; i < 26; i++ {
		res += bitCount(arr[i])
	}
	return res
}
