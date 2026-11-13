package novemember

func longestMonotonicSubarray(nums []int) int {
	ans := 0
	count := 0
	n := len(nums)
	if nums[1] > nums[0] {
		count = 1
	}
	for i := 2; i < n; i++ {
		if nums[i] > nums[i-1] {
			count++
		} else {
			if count > ans {
				ans = count
				count = 0
			}
		}
	}
	if nums[1] < nums[0] {
		count = 1
	}
	for i := 2; i < n; i++ {
		if nums[i] < nums[i-1] {
			count++
		} else {
			if count > ans {
				ans = count
				count = 0
			}
		}
	}
	return ans
}
