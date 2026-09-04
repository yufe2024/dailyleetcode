package new

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func minimumDeletions(nums []int) int {
	n := len(nums)
	min := nums[0]
	max := nums[0]
	x := 0
	y := 0
	for i := 1; i < n; i++ {
		if nums[i] < min {
			min = nums[i]
			x = i
		}
		if nums[i] > max {
			max = nums[i]
			y = i
		}
	}
	qq := max(x, y) + 1
	if x > y {
		x, y = y, x
	}
	ans := qq
	qh := n - y + x
	ans = max(qh, ans)
	hh := n - x
	ans := max(hh, ans)
	return ans
}
