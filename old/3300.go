package old

func solve(a int) int {
	ans := 0
	for a/10 != 0 {
		ans += a % 10
		a /= 10
	}
	ans += a
	return ans
}
func minElement(nums []int) int {
	ans := solve(nums[0])
	for i := 0; i < len(nums); i++ {
		if solve(nums[i]) <= ans {
			ans = solve(nums[i])
		}
	}

	return ans
}
