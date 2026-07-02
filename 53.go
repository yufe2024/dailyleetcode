package main

func maxSubArray(nums []int) int {
	an := nums[0]
	var ans []int
	ans = append(ans, max(nums[0], 0))
	for i := 1; i < len(nums); i++ {
		ans = append(ans, max(ans[i-1]+nums[i], nums[i]))
		if ans[i] > an {
			an = ans[i]
		}
	}
	return an
}
