package main

func subarraySum(nums []int, k int) int {
	ans := 0
	cur := 0
	m := make(map[int]int)
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		cur += nums[i]
		if _, ok := m[cur-k]; ok {
			ans += m[cur-k]
		}
		m[cur]++
	}
	return ans
}
