package main

func diStringMatch(s string)[]int{
	n := len(s)
	left , right := 0 , n
	ans := make([]int,n+1)
	for i,v := range s{
		if v == 'I'{
			ans[i] = left
			left++
		}else{
			ans[i] = right
			right--
		}
	}
	ans[n] = left
	return ans	
}