package main

func buddyStrings(s string, goal string) bool {
   
    if len(s) != len(goal) {
        return false
    }

    if s == goal {
        freq := make(map[rune]int)
        for _, c := range s {
            freq[c]++
            if freq[c] >= 2 {
                return true
            }
        }
        return false
    }

    diff := make([]int, 0)
    for i := 0; i < len(s); i++ {
        if s[i] != goal[i] {
            diff = append(diff, i)
        }
    }

    if len(diff) != 2 {
        return false
    }

    return s[diff[0]] == goal[diff[1]] && s[diff[1]] == goal[diff[0]]
}