package october

func removeOuterParentheses(s string) string {
	n := len(s)
	if n == 0 {
		return s
	}
	left := 0
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			left++
		}
	}
}
