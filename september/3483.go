package september

// 简单题吗，难道我真的是个弱智
func totalNumbers(digits []int) int {
	countji := 0
	countou := 0
	have0 := 0
	num := make([]int, 11)
	for _, d := range digits {
		num[d]++
	}
	for i := 0; i <= 9; i++ {
		if i == 0 && num[i] > 0 {
			have0 = 1
		} else if num[i] > 0 && num[i]%2 == 0 {
			countou++
			num[i] = 0
		} else if num[i] > 0 && num[i]%2 == 1 {
			countji++
			num[i] = 0
		}
	}
	if coutou+coutji < 3 {
		return 0
	}
	if countji+countou == 3 {
		if have0 == 0 {
			//2奇1偶数 2
			//1奇数 2偶数 4
			//3偶数 6
			return countou * 2
		} else {
			if countou == 0 {
				return 2
			}
			if countou == 1 {
				return 3
			}
			if countou == 2 {
				return 4
			}
		}
	}

	if have0 == 0 {
		return (countji + countou - 1) * (countji + countou - 2) * countou
	} else {
		return (countji+countou-2)*(countji+countou-3)*countou + (countji+countou-1)*(countji+countou-2)
	}

}

//题解写的挺好的，枚举就可以，标记一下是否visit
