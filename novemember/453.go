package novemember

import "sort"

func minMoves(nums []int) int {
	//最大的数一开始不加，直到出现两个最大的数，然后每操作两次来平衡数组，有几个就要操作几次
	//好吧，所有数字加一等于找数字减一，所以我讨厌贪心
	sort.Ints(nums)
	min := nums[0]
	count := 0
	for i := 1; i < len(nums); i++ {
		count += nums[i] - min
	}
	return count
}
