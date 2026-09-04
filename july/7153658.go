package july

func gcd(x, y int) int {
	if x < y {
		x, y = y, x
	}
	for y != 0 {
		x, y = y, x%y
	}
	return x

}
func gcdOfOddEvenSums(n int) int {
	a := 2 * n
	return gcd((a)*n/2, (a)*n/2+n)
}
