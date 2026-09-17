package july

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func pathExistenceQueries(n int, nums []int, maxDiff int, queries [][]int) []bool {
	ans := make([][]bool, n+1)
	for i := range ans {
		ans[i] = make([]bool, n+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				ans[i][j] = true
			} else if abs(nums[i]-nums[j]) > maxDiff {
				ans[i][j] = true
			} else {
				ans[i][j] = false
			}
		}
	}
	answer := make([]bool, n)
	for i := 0; i < n; i++ {
		x := queries[i][0]
		y := queries[i][1]
		if ans[x][y] {
			answer[i] = true
		} else {
			answer[i] = false
		}
	}
	return answer
}
