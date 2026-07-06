package old

func longestConsecutive(nums []int) int {
	ans := 0
	mp := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		mp[nums[i]]++
	}
	count := 0
	for i, _ := range mp {
		if mp[i-1] > 0 {
			continue
		} else {
			count++
			for j := i; ; j++ {
				if mp[j] > 0 {
					count++
				} else {
					break
				}
			}
		}
		if count > ans {
			ans = count

		}
		count = 0
	}
	return ans
}
