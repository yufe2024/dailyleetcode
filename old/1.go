package old

import "sort"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	ans := make(map[int]int)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] == 1 {
			return []int{i, ans[target-nums[i]]}
		} else {
			m[target-nums[i]] = 1
			ans[target-nums[i]] = i
		}
	}
	return []int{-1, -1}
}
