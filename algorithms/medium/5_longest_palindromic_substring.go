package medium

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := pattern(s, i, i)
		left2, right2 := pattern(s, i, i+1)
		if right1-left1 > right-left {
			left, right = left1, right1
		}
		if right2-left2 > right-left {
			left, right = left2, right2
		}
	}
	return s[left : right+1]
}

func pattern(s string, i, j int) (int, int) {
	for ; i >= 0 && j < len(s) && s[i] == s[j]; i, j = i-1, j+1 {
	}
	return i + 1, j - 1
}
