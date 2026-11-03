package october

import "sort"

func maxAlternatingSum(nums []int) int64 {
	n := len(nums)
	sort.Ints(nums)
	num2 := make([]int64, n)
	for i := 0; i < len(nums); i++ {
		num2[i] = int64(nums[i]) * int64(nums[i])
	}

	var sum int64
	for i := 0; i < len(nums)/2; i++ {
		sum += num2[i]
	}
	var sum2 int64
	for i := len(nums) / 2; i < len(num2); i++ {
		sum2 += num2[i]
	}
	return sum2 - sum
}
