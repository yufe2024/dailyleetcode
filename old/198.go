package old

func rob(nums []int) int {
	n := len(nums)
	ans := make([]int, n)
	if n == 1 {
		return nums[0]
	}
	ans[0] = nums[0]
	ans[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		ans[i] = max(ans[i-2]+nums[i], ans[i-1])
	}
	return ans[len(nums)-1]
}
