package main

func canJump(nums []int) bool {
	ml := 0
	for i := 0; i < len(nums); i++ {
		if ml >= i {
			ml = max(ml, nums[i]+i+1)
		}

	}
	if ml > len(nums) {
		return true
	}
	return false
}
