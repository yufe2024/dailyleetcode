package september

func firstStableIndex(nums []int, k int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	ma := make([]int, n)
	mi := make([]int, n)
	ma[0] = nums[0]
	mi[n-1] = nums[n-1]
	for i := 1; i < n; i++ {
		if nums[i] > ma[i-1] {
			ma[i] = nums[i]
		} else {
			ma[i] = ma[i-1]
		}
	}
	for i := n - 2; i >= 0; i-- {
		if nums[i] < mi[i+1] {
			mi[i] = nums[i]
		} else {
			mi[i] = mi[i+1]
		}
	}
	for i := 0; i < n; i++ {
		if ma[i]-mi[i] <= k {
			return i
		}
	}
	return -1
}
