package novemember

import "sort"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func minMoves2(nums []int) int {
	sort.Ints(nums)
	ans := 0
	num := nums[len(nums)/2]
	for i := 0; i < len(nums); i++ {
		ans += abs(num - nums[i])
	}
	return ans
}
