package strings

import "fmt"

//参考strings count函数
func RabinKarp(src, search string) {
	// Rabin-Karp search
	hashSearch, pow := hashStr(search)
	matchPostinMap := make([]int, 0)
	h := uint32(0)
	for i := 0; i < len(search); i++ {
		h = h*primeRK + uint32(src[i])
	}
	if h == hashSearch && src[:len(search)] == search {
		matchPostinMap = append(matchPostinMap, 0)
		// lastmatch = len(search)
	}
	for i := len(search); i < len(src); {
		h *= primeRK
		h += uint32(src[i])
		h -= pow * uint32(src[i-len(search)])
		i++
		if h == hashSearch && src[i-len(search):i] == search {
			matchPostinMap = append(matchPostinMap, i-len(search))
		}
	}
	fmt.Println(matchPostinMap)
}

// hashStr returns the hash and the appropriate multiplicative
// factor for use in Rabin-Karp algorithm.
//模仿strings.Count需要先计算字符串hash值并且返回加减首位应该乘的值
const primeRK = 16777619

func hashStr(search string) (uint32, uint32) {
	hash := uint32(0)
	for i := 0; i < len(search); i++ {
		hash = hash*primeRK + uint32(search[i])
	}
	var pow, sq uint32 = 1, primeRK
	for i := len(search); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}
