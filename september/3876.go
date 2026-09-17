package september

import "sort"

func uniformArray(nums1 []int) bool {
	//先按奇数的情况查一遍，再按偶数的情况查一遍
	//考虑一下什么情况会不过，2，3就不行，为什么，因为2比较小，只能以它为标准，3怎么搞都不会和2相同，那好做了吧应该
	n := len(nums1)
	if n == 0 {
		return true
	}
	sort.Ints(nums1)
	flag := -1
	if nums1[0]%2 == 0 {
		flag = 0
	} else {
		flag = 1
	}
	havej := 0
	if flag == 0 {
	} else {
		havej = 1
	}

	for i := 1; i < len(nums1)-1; i++ {
		if flag == 1 {
			if nums[i]%2 == 0 && havej == 0 {
				return false
			}
			if nums1[i]%2 == 1 {
				havej = 1
			}
		} else {
			if nums1[i]%2 == 1 {
				return false
			}

		}
	}
	return true
}
