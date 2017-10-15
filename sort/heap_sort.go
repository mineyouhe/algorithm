package sort

import (
	"fmt"
)

func AdjustHeap(ar []int, index int) {
	left := (index+1)*2 - 1
	right := (index + 1) * 2
	iMax := index
	n := len(ar)
	// fmt.Println(n, index, left, right)
	if left < n && ar[left] > ar[iMax] {
		iMax = left
	}
	if right < n && ar[right] > ar[iMax] {
		iMax = right
	}
	// fmt.Println(iMax)
	if index != iMax {
		Swap(ar, index, iMax)
		AdjustHeap(ar, iMax)
		// fmt.Println(ar)
	}
}

func Swap(ar []int, pre, after int) {
	ar[pre], ar[after] = ar[after], ar[pre]
}

//最大堆
func HeapSort(ar []int) {
	n := len(ar)
	lastParent := (n - 1) / 2
	for i := lastParent; i >= 0; i-- {
		AdjustHeap(ar, i)
	}
	for i := n; i > 0; i-- {
		Swap(ar, 0, i-1)
		AdjustHeap(ar[0:i-1], 0)
	}
	fmt.Println(ar)
}
