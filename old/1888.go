package old

func minFlips(s string) int {
	sr := s + s
	lens := len(s)
	ans := 0
	count1 := 0
	count2 := 0
	for i := 0; i < lens; i++ {
		if i%2 == 0 {
			if sr[i] == '0' {
				count1++
			} else {
				count2++
			}
		}
		if i%2 != 0 {
			if sr[i] == '1' {
				count1++
			} else {
				count2++
			}
		}
		ans = count1
		if count1 > count2 {
			ans = count2
		}
	}
	for j := 0; j < lens-1; j++ {
		if j%2 == 0 {
			if sr[j] != '1' {
				count1--
			}
			if sr[j] != '0' {
				count2--
			}
		} else {
			if sr[j] != '0' {
				count1--
			}
			if sr[j] != '1' {
				count2--
			}
		}
		if (j+lens)%2 == 0 {
			if sr[j+lens] != '1' {
				count1++
			}
			if sr[j+lens] != '0' {
				count2++
			}
		} else {
			if sr[j+lens] != '0' {
				count1++
			}
			if sr[j+lens] != '1' {
				count2++
			}
		}
		if count1 < ans {
			ans = count1
		}
		if count2 < ans {
			ans = count2
		}
	}
	return ans
}
