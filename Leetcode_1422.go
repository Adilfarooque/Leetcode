package main


func maxScore(s string)int {
    n := len(s)
	score := 0
	if s[0] == '0' {
		score++
	}
	for i := 1; i < n; i++ {
		if s[i] == '1' {
			score++
		}
	}
	maxScore := score
	for i := 1; i < n-1; i++ {
		if s[i] == '0' {
			score++
		}else{
			score--
		}
		if score > maxScore {
			maxScore = score
		}
	}
	return maxScore
}
