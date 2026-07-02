package main

import (
	"slices"
)

func merge(intervals [][]int) [][]int {
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	var ans [][]int
	ans = append(ans, intervals[0])
	for i := 1; i < len(intervals); i++ {
		last := ans[len(ans)-1]
		current := intervals[i]
		if last[1] >= current[0] {
			last[1] = max(last[1], current[1])
		} else {
			ans = append(ans, current)
		}
	}
	return ans
}
