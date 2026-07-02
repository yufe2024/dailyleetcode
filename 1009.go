package main

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}
	temp := n
	var a []int
	var ans int
	for i := 0; temp != 0; i++ {
		a = append(a, temp%2)
		temp = temp / 2
	}
	for i := 0; i < len(a); i++ {
		if a[i] == 1 {
			a[i] = 0
		} else {
			a[i] = 1
		}
	}
	for i, j := 0, len(a)-1; i < j; i++ {
		a[i], a[j] = a[j], a[i]
		j--
	}
	j := 1
	for i := len(a) - 1; i >= 0; i-- {
		ans += a[i] * j
		j = j * 2
	}
	return ans
}
