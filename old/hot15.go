package old

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	n := len(nums)
	for i := 0; i < n-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		k := n - 1
		for j := i + 1; j < k; {
			if nums[i]+nums[j]+nums[k] == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				k--
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if nums[i]+nums[j]+nums[k] > 0 {
				k--

			} else {
				j++

			}
		}
	}

	return ans
}
