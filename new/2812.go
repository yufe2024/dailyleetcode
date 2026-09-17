package new

func maximumSafenessFactor(grid [][]int) int {
	dirs := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	n := len(grid)
	dis := make([][]int, n)
	for i := range dis {
		dis[i] = make([]int, n)
		for j := range dis[i] {
			if grid[i][j] == 1 {
				dis[i][j] = 1
			} else {
				dis[i][j] = -1
			}
		}
	}
	q := make([][]int, 0)
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				q = append(q, []int{i, j})
			}
		}
	}
	for len(q) > 0 {
		cx, cy := q[0][0], q[0][1]
		q = q[1:]
		for _, dir := range dirs {
			nx := cx + dir[0]
			ny := cy + dir[1]
			if nx >= 0 && nx < n && ny >= 0 && ny < n && dis[nx][ny] == -1 {
				q = append(q, []int{nx, ny})
				dis[nx][ny] = dis[cx][cy] + 1
			}
		}
	}
	check := func(limit int) bool {
		visited := make([][]bool, n)
		for i := range visited {
			visited[i] = make([]bool, n)
		}
		q := [][]int{{0, 0}}
		visited[0][0] = true
		for len(q) > 0 {
			cx, cy := q[0][0], q[0][1]
			q = q[1:]
			for _, dir := range dirs {
				nx := cx + dir[0]
				ny := cy + dir[1]
				if nx < 0 || nx >= n || ny < 0 || ny >= n || visited[nx][ny] || dis[nx][ny] < limit {
					continue
				}
				q = append(q, []int{nx, ny})
				if nx == n-1 && ny == n-1 {
					return true
				}
			}

		}
		return false
	}
	right := 2 * n
	left := 0
	for left != right {
		mid := (left + right) / 2
		if check(mid) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}
