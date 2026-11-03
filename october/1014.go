package october

func maxSumOfSquares(num int, sum int) string {
	count := 0
	for sum >= 9 {
		sum -= 9
		count++
	}
	check := 0
	if sum > 0 {
		count++
		check = 1

	}
	if count >= num {
		return ""
	}
	ans := make([]byte, num)
	for i := 0; i < count; i++ {
		ans[i] = '9'
	}
	if check == 1 {
		ans[count-1] = byte(sum + int('0'))
	}

	for i := count; i < num; i++ {
		ans[i] = '0'
	}
	return string(ans)
}
