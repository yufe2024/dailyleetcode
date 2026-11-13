package novemember

func minimumCost(cost1 int, cost2 int, costBoth int, need1 int, need2 int) int64 {
	if need1 > need2 {
		need1, need2 = need2, need1
		cost1, cost2 = cost2, cost1
	}
	if costBoth < cost2 {
		return int64(costBoth) * int64(need2)
	}
	if costBoth > cost1+cost2 {
		return int64(cost1)*int64(need1) + int64(cost2)*int64(need2)
	}
	return int64(costBoth)*int64(need1) + int64(cost2)*int64(need2-need1)
}
