package old

func check(time int, workerTimes []int, mountainHeight int) bool {
	hei := 0
	for i := 0; i < len(workerTimes); i++ {
		x := 0
		left := 0
		right := mountainHeight
		for left <= right {
			mid := left + (right-left)/2

			if (mid*mid+mid)*workerTimes[i] < 2*time {
				x = mid
				right = mid - 1
			} else {
				left = mid + 1
			}

		}
		hei += x
	}
	if hei < mountainHeight {
		return false
	} else {
		return true
	}
}

func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	ans := 0
	left := 0
	right := 1000000000
	for left <= right {
		mid := left + (right-left)/2
		if check(mid, workerTimes, mountainHeight) {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return int64(ans)
}
