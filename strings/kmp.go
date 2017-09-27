package strings

import "fmt"

//KMP算法,先计算search字符串部分匹配表,然后开始比较,不相同则根据部分匹配表移位
func KMP(src, search string) {
	postionInfo := MoveInfoTable(search)
	searchIndex := 0
	srcIndex := 0
	searchLen := len(search)
	srcLen := len(src)
	matchPostinMap := make([]int, 0)
	for srcIndex < srcLen {
		// fmt.Println(srcIndex, searchIndex)
		if src[srcIndex] == search[searchIndex] {
			searchIndex += 1
			srcIndex += 1
			// fmt.Println(srcIndex, searchIndex, searchLen)
			if searchIndex == searchLen {
				searchIndex = 0
				matchPostinMap = append(matchPostinMap, srcIndex-searchLen)
				srcIndex = srcIndex - searchLen + 1
			}
		} else {
			if searchIndex == 0 {
				srcIndex += 1
				continue
			} else {
				// moveLen := searchIndex + 1 - postionInfo[searchIndex-1]
				// searchIndex -= moveLen
				searchIndex = postionInfo[searchIndex-1]
			}
		}
	}
	fmt.Println(matchPostinMap)
}

//搜索字符串移动信息表
func MoveInfoTable(search string) map[int]int {
	position := make(map[int]int)
	for k, _ := range search {
		prefixS := search[0:k]
		suffixS := search[1 : k+1]
		prefixM := make(map[string]int)
		for kk, _ := range prefixS {
			prefixM[prefixS[0:kk+1]] = kk + 1
		}
		maxLen := 0
		// fmt.Println(prefixS, suffixS)
		for kk, _ := range suffixS {
			if l, ok := prefixM[suffixS[kk:]]; ok {
				if l > maxLen {
					maxLen = l
					// fmt.Println(prefixM, suffixS[kk:], maxLen)
				}
			}
		}
		position[k] = maxLen
	}
	// fmt.Println(position)
	return position
}
