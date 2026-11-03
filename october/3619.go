package october

func countIslands(grid [][]int, k int) int {
	//回溯，找到陆地然后四周入队列，加到不能再加，找下一片没有visit的陆地
	dir := make([][]int, 4)
	m := len(grid)
	n := len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {

			}
			visited[i][j] = true
		}
	}
}
