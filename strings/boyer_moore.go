package strings

import "strings"

//预先计算好字符跳转表，坏字符跳转表，只与search字符相关，与其他无关
func BoyerMoore(src, search string) int {
	var badChar [256]int
	goodSuffixSkip := make([]int, len(search))
	last := len(search) - 1
	//构造坏字符数组
	for i := range badChar {
		badChar[i] = len(search)
	}
	for i := 0; i < last; i++ {
		badChar[search[i]] = last - i
	}
	//goodSuffixSkip 存储应该后移的长度
	//(2).对于Good-suffix后移位数计算方法为：
	// 1.pattern中间的某一子字符串pattern[j-s+1:m-s] == pattern[j+1:m]，可将pattern右移s位；
	// 2.pattern已比较部分pattern[j+1:m]的后缀pattern[s+1:m]与pattern的前缀pattern[1:m-s]相同，可将pattern右移s位。
	// 满足上面情况的s的最小值为最佳右移距离。

	lastPrefix := last
	for i := last; i >= 0; i-- {
		if strings.HasPrefix(search, search[i+1:]) {
			lastPrefix = i + 1
			//长度
		}
		//如果匹配到，则为最后一个位置，如果匹配不到
		goodSuffixSkip[i] = lastPrefix + last - i
		// fmt.Println(lastPrefix, i, goodSuffixSkip[i])
		// lastPrefix is the shift, and (last-i) is len(suffix).
		//后缀的最后一位位置
	}
	// fmt.Println(goodSuffixSkip, lastPrefix)
	for i := 0; i < last; i++ {
		lenSuffix := longestCommonSuffix(search, search[1:i+1])
		// fmt.Println(lenSuffix, i, search[1:i+1], search[i-lenSuffix], search[last-lenSuffix])
		if search[i-lenSuffix] != search[last-lenSuffix] {
			// (last-i) is the shift, and lenSuffix is len(suffix).
			goodSuffixSkip[last-lenSuffix] = lenSuffix
			// fmt.Println(last-lenSuffix, goodSuffixSkip[last-lenSuffix], i, last, lenSuffix, search[1:i+1])
		}
	}
	// fmt.Println(goodSuffixSkip)
	i := len(search) - 1
	for i < len(src) {
		// Compare backwards from the end until the first unmatching character.
		j := len(search) - 1
		for j >= 0 && src[i] == search[j] {
			i--
			j--
		}
		if j < 0 {
			return i + 1 // match
		}
		// fmt.Println(max(badChar[src[i]], goodSuffixSkip[j]))
		i += max(badChar[src[i]], goodSuffixSkip[j])
	}
	return -1
}

//获取最长公共后缀
func longestCommonSuffix(a, b string) (i int) {
	for ; i < len(a) && i < len(b); i++ {
		if a[len(a)-1-i] != b[len(b)-1-i] {
			break
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
