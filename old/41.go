package old

func firstMissingPositive(nums []int) int {
	n := len(nums)
	s := make([]int, n+1)
	for i := 0; i < n; i++ {
		if nums[i] > n || nums[i] <= 0 {
			continue
		} else {
			s[nums[i]] = 1
		}
	}
	for i := 1; i < n; i++ {
		if s[i] != 0 {
			return i
		}
	}
	return 1
}
