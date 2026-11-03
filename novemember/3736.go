package novemember

func minMoves(nums []int) int {
	ans := 0
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		ans += max - nums[i]
	}
	return ans
}
