package july

func removeCoveredIntervals(intervals [][]int) int {
	ans := len(intervals)
	for i := 0; i < len(intervals); i++ {
		for j := 0; j < len(intervals); j++ {
			if i != j {
				if intervals[i][0] >= intervals[j][0] && intervals[i][1] <= intervals[j][1] {
					ans--
					break
				}
			}

		}
	}
	return ans
}
