// package main
//
//	func maxProduct(nums []int) int {
//		ans := 0
//		n := len(nums)
//		la := -1
//		dp := make([]int, n)
//		dp[0] = nums[0]
//		for i := 1; i < n; i++ {
//			if nums[i] > 0 {
//				dp[i] = dp[i-1] * nums[i]
//				if dp[i] > ans {
//					ans = dp[i]
//				}
//			} else if nums[i] == 0 {
//				la = -1
//				if dp[i] > ans {
//					ans = dp[i]
//				}
//				dp[i] = 0
//			} else {
//				if la == -1 {
//					la = i
//					continue
//				} else {
//					dp[i] = (dp[i-1] * nums[i]) / dp[la-1]
//					if dp[i] > ans {
//						ans = dp[i]
//					}
//				}
//			}
//		}
//		return ans
//	}
package old

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 初始化最大值、最小值以及全局最终答案为第一个元素
	curMax := nums[0]
	curMin := nums[0]
	ans := nums[0]

	for i := 1; i < len(nums); i++ {
		x := nums[i]

		// 如果当前数是负数，乘以它之后，最大值会变成最小值，最小值会变成最大值
		// 所以我们提前交换 curMax 和 curMin
		if x < 0 {
			curMax, curMin = curMin, curMax
		}

		// 更新当前位置的最大乘积和最小乘积
		curMax = max(x, curMax*x)
		curMin = min(x, curMin*x)

		// 维护全局最大值
		if curMax > ans {
			ans = curMax
		}
	}

	return ans
}
