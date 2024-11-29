package main

func backspaceCompare(s string, t string) bool {
	var stack1, stack2 []rune
	for _, c := range s {
		if c != '#' {
			stack1 = append(stack1, c)
		} else if len(stack1) > 0 {
			stack1 = stack1[:len(stack1)-1]
		}
	}
	for _, c := range t {
		if c != '#' {
			stack2 = append(stack2, c)
		} else if len(stack2) > 0 {
			stack2 = stack2[:len(stack2)-1]
		}
	}
	return string(stack1) == string(stack2)
}
