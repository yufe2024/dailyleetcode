package main

func maxSlidingWindow(nums []int, k int) []int {
	left := 0
	right := k - 1
	var ans []int
	m := nums[0]
	for i := left; i <= right; i++ {
		if nums[i] > m {
			m = nums[i]
		}
	}
	ans = append(ans, m)
	for right < len(nums)-1 {
		right++
		left++
		if nums[right] > ans[right-1] && nums[left] != ans[left-1] {
			ans = append(ans, nums[right])
		} else {
			ans = append(ans, ans[left]-1)
		}
	}
	return ans
}
