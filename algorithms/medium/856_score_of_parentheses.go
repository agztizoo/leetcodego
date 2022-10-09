package medium

func scoreOfParentheses(s string) int {
	value, bal := 0, 0
	for i, x := range s {
		if x == '(' {
			bal++
		} else {
			bal--
			if s[i-1] == '(' {
				value += 1 << bal
			}
		}
	}
	return value
}
