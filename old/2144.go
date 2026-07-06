package old

import "sort"

func minimumCost(cost []int) int {
	sort.Ints(cost)
	ans := 0
	n := len(cost)
	for i := n - 1; i >= 0; {
		if i == 0 {
			ans += cost[i]
			return ans
		} else if i <= 2 {
			ans += cost[i]
			ans += cost[i-1]
			return ans
		}

		ans += cost[i] + cost[i-1]
		i = i - 3
	}
	return ans
}
