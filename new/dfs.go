package new

func dfs(grid [][]int) {
	used := make(map[int]bool)
	q := [][]int{}
	n := len(grid)
	loop := func() {
		if len(q) == n {
			return
		}
		for _, i := range q {
			if !used[i] {
				q = append(q, i)
				used[i] = true
				loop()
				used[i] = false
				q = q[:len(q)-1]
			} else {
				continue
			}

		}
	}

}
