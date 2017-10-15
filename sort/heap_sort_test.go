package sort

import (
	"testing"
)

func Test_heap_sort(t *testing.T) {
	ar := []int{1, 9, 78, 343, 312321, 3, 4, -8, 21312321321, 323}
	HeapSort(ar)
}
