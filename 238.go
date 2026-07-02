package main

func productExceptSelf(nums []int) []int {
	pre := make([]int, len(nums))
	pre[0] = 1
	last := make([]int, len(nums))
	last[len(nums)-1] = 1
	for i := 1; i < len(nums); i++ {
		pre[i] = nums[i-1] * pre[i-1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		last[i] = nums[i+1] * last[i+1]
	}
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = pre[i] * last[i]
	}
	return ans
}
