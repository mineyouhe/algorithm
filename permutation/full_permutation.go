package permutation

import (
	"fmt"
	"strings"
)

// func main() {
// 	tmpSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
// 	m := make(map[string]int)
// 	tmpFunc(tmpSlice, 0, m, 5)
// 	for k, v := range m {
// 		fmt.Println(k, v)
// 	}
// }
//n个元素中选择m个进行全排列
func partition_full_permutation(t []int, start int, findM map[string]int, m int) {
	if start == len(t)-1 {
		key := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(t[0:m])), ","), "[]")
		if _, ok := findM[key]; !ok {
			findM[key] = 1
		} else {
			findM[key] = findM[key] + 1
		}
	} else {
		for i := start; i < len(t); i++ {
			tmp0, tmp1 := t[start], t[i]
			t[start], t[i] = t[i], t[start]
			partition_full_permutation(t, start+1, findM, m)
			t[start], t[i] = tmp0, tmp1
		}
	}
}

//全排列
func full_permutation(t []int, start int) {
	if start == len(t)-1 {
		fmt.Println(t)
	} else {
		for i := start; i < len(t); i++ {
			tmp0, tmp1 := t[start], t[i]
			t[start], t[i] = t[i], t[start]
			full_permutation(t, start+1)
			t[start], t[i] = tmp0, tmp1
		}
	}
}
