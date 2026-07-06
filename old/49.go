package old

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	mapp := make(map[string][]string)
	for _, str := range strs {
		b := []byte(str)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		mapp[string(b)] = append(mapp[string(b)], str)
	}
	ans := make([][]string, 0)
	for _, str := range mapp {
		ans = append(ans, str)
	}
	return ans
}
