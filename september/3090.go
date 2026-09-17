package september

func firstStableIndex(nums []int, k int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	// 前缀最大值
	prefMax := make([]int, n)
	prefMax[0] = nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > prefMax[i-1] {
			prefMax[i] = nums[i]
		} else {
			prefMax[i] = prefMax[i-1]
		}
	}
	// 后缀最小值（从 i 到末尾）
	sufMin := make([]int, n)
	sufMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		if nums[i] < sufMin[i+1] {
			sufMin[i] = nums[i]
		} else {
			sufMin[i] = sufMin[i+1]
		}
	}
	// 查找第一个满足条件的下标
	for i := 0; i < n; i++ {
		if abs(prefMax[i]-sufMin[i]) <= k {
			return i
		}
	}
	return -1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
