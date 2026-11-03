package new

func checkValidString(s string) bool {
	left := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else if s[i] == '*' {
			if i == len(s)-1 {
				if i==1||i==0{
					return false
				}
			}else{
				
			}

		} else if s[i] == ')' {
			if left > 0 {
				left--
			}
		}
	}
}
