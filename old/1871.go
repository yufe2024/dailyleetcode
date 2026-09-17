package old

func canReach(s string, minJump int, maxJump int) bool {
	num := []byte(s)
	n := len(num)
	f := make([]int, n)
	var num2 []int
	for i := 0; i < n; i++ {
		if num[i] == 0 {
			num2 = append(num2, i)
		}
	}
	f[0] = 1
	n2 := len(num2) - 1
	for i := 1; i < n; i++ {
		if num[i] == 1 {
			f[i] = 0
		} else if {

		}
	}
}
