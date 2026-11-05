package novemember

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1
	for left < right {
		mid := left + (right-left)/2
		if mid != len(arr)-1 {
			if arr[mid] < arr[mid+1] {
				left = mid + 1
			} else {
				right = mid
			}
		} else {
			return mid
		}

	}
	return left
}
