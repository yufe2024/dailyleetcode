package novemember

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	ans := 0
	n := len(arr1)
	m := len(arr2)
	for i := 0; i < n; i++ {
		check := 0
		for j := 0; j < m; j++ {
			if abs(arr1[i]-arr2[j]) <= d {
				check = 1
			}
		}
		if check == 0 {
			ans++
		}
	}
	return ans
}
