package novemember

func pinjie(x, y int) int {
	ans := x

	nums := make([]int, 0)
	for y != 0 {
		nums = append(nums, y%10)
		y /= 10
	}
	for i := len(nums) - 1; i >= 0; i-- {
		ans = ans*10 + nums[i]
	}
	return ans
}
func minMergedInteger(nums []int) int {
	min := pinjie(nums[0], nums[1])
	for i := 0; i < len(nums); {
		if pinjie(nums[i], nums[i+1]) <= min {
			min = pinjie(nums[i], nums[i+1])
		}
		i = i + 2
	}
	return min
}
