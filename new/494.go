package new

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	if (sum+target)%2 == 1 {
		return 0
	}
	n := len(nums)
	var dfs func(nums []int, i int, c int) int
	dfs = func(nums []int, i int, c int) int {
		if c == 0 {
			return 1
		}
		if i < n-1 && c > 0 {
			return dfs(nums, i+1, c) + dfs(nums, i+1, c-nums[i])
		}
		return 0
	}
	return dfs(nums, 0, (sum+target)/2)
}
