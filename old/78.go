package old

func subsets(nums []int) [][]int {
	var ans [][]int
	var path []int
	var dfs func(start int)
	dfs = func(start int) {
		temp := make([]int, len(path))
		copy(temp, path)
		ans = append(ans, temp)
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])

			dfs(i + 1)

			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}
