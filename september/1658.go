package september

func minOperations(nums []int, x int) int {
	ans := 100000007
	n := len(nums)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	solve := func(left, right int   ) int {}

}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
