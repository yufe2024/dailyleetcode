package novemember

import "sort"

func countServers(n int, logs [][]int, x int, queries []int) []int {
	m := len(queries)
	numLogs := len(logs)

	sort.Slice(logs, func(i, j int) bool {
		return logs[i][1] < logs[j][1]
	})

	type queryItem struct {
		time int
		idx  int
	}
	sorted := make([]queryItem, m)
	for i, q := range queries {
		sorted[i] = queryItem{time: q, idx: i}
	}

	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].time < sorted[j].time
	})

	ans := make([]int, m)
	cnt := make([]int, n+1)
	count := 0
	left, right := 0, 0

	for _, q := range sorted {
		qTime := q.time
		leftBound := qTime - x

		for right < numLogs && logs[right][1] <= qTime {
			sid := logs[right][0]
			if cnt[sid] == 0 {
				count++
			}
			cnt[sid]++
			right++
		}

		for left < numLogs && logs[left][1] < leftBound {
			sid := logs[left][0]
			cnt[sid]--
			if cnt[sid] == 0 {
				count--
			}
			left++
		}

		ans[q.idx] = n - count
	}

	return ans
}