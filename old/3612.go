package old

func isLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}
func processStr(s string) string {
	nums := []byte(s)
	for _, a := range s {
		if a == '*' {
			n := len(nums)
			for {
				if isLower(nums[n-1]) {
					nums = nums[:n-1]
				} else {
					n--
				}
				if n == 0 {
					break
				}
			}
		}
		if a == '#' {
			for _, b := range nums {
				if isLower(byte(b)) {
					nums = append(nums, b)
				}
			}
		}
		if a == '%' {
			i := 0
			j := len(s) - 1
			for i < j {
				if !isLower(nums[i]) {
					i++
				}
				if !isLower(nums[j]) {
					j--
				}
				if isLower(nums[j]) && isLower(nums[i]) {
					nums[i], nums[j] = nums[j], nums[i]
				}

			}
		}
	}
	return string(nums)
}
