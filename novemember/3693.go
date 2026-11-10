package novemember

func minn(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
}

//	func min(a, b int) int {
//		if a <= b {
//			return a
//		}
//
//		return b
//	}
func climbStairs(n int, costs []int) int {
	ans := make([]int, n)
	ans[0] = costs[0] + 1
	if n == 0 {
		return ans[0]
	}
	if n > 1 {
		ans[1] = min(ans[0]+1, 4) + costs[1]
	} else {
		return 0
	}
	if n > 2 {
		ans[2] = costs[2] + minn(ans[0]+4, ans[1]+1, 9)
	} else {
		return ans[1]
	}
	for i := 3; i < n; i++ {
		ans[i] = minn(ans[i-3]+9, ans[i-2]+4, ans[i-1]+1) + costs[i]
	}
	return ans[n-1]
}
