package main

import "sort"

func asteroidsDestroyed(mass int, asteroids []int) bool {
	var ans bool
	sort.Ints(asteroids)
	for i := 0; i < len(asteroids); i++ {
		if mass > asteroids[i] {
			mass += asteroids[i]
		} else {
			ans = false
			return ans
		}
	}
	return true
}
