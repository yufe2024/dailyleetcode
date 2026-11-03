package novemember

func minimumMaxValue(n int, index int, value int) int64 {
	//从index开始看，往后看为了找min，往前看判断是否合法
	nums := make([]int, n)
	nums[index] = value
	//必须不是素数，素数也非法，一个一个往前找，前一个数字尽可能大， --再看有没有是傻子的做法
	//往前有点复杂，因为是可以即可，没必要找最大的，可以考虑从开头开始235711这样尝试能不能走到最后

}
