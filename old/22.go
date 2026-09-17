package old

func generateParenthesis(n int) []string {
	var ans []string
	var path []byte
	var dfs func(l int, r int)
	dfs = func(l int, r int) {
		if len(path) == 2*n {
			ans = append(ans, string(path))
		}
		if l < n {
			path = append(path, '(')
			dfs(l+1, r)
			path = path[:len(path)-1]
		}
		if r < l {
			path = append(path, ')')
			dfs(l, r+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)
	return ans
}
