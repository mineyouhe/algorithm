package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
	//fmt.Println(Bubble([]int{12345, 7, 23, 123, 1, 2, 3, 7666}))
	//fmt.Println(Insert([]int{12345, 7, 23, 123, 1, 2, 3, 7666}))
	//fmt.Println(Selection([]int{12345, 7, 23, 123, 1, 2, 3, 7666}))
	//	fmt.Println(QuickSort([]int{12345, 7, 23, 123, 1, 2, 3, 7666}))
	//fmt.Println(MegerSort([]int{12345, 7, 23, 123, 1, 2, 3, 7666}))
	//fmt.Println(HeapSort([]int{12345, 7, 23, 123, 1, 2, 3, 7666}))
	//fmt.Println(RadixSort([]int{12345, 7, 23, 123, 1, 2, 3, 7666}))
	//fmt.Println(CountSort([]int{12345, 7, 23, 123, 1, 2, 3, 7666}))
	//fmt.Println(BucketSort([]int{1234, 7, 23, 123, 1, 2, 3, 666}))
	fmt.Println(ShellSort([]int{1234, 7, 23, 123, 1, 2, 3, 666}))
}

//冒泡
//两两比较，前比后大的交换
//O(N^2)
func Bubble(s []int) []int {
	t := len(s)
	for i := 0; i < t-1; i++ {
		for j := 1; j < t-i; j++ {
			if s[j-1] > s[j] {
				s[j-1], s[j] = s[j], s[j-1]
			}
		}
	}
	return s
}

//插入
//默认当前子序列有序，依次加入数字，向前寻找自己位置
//O(N^2)
func Insert(s []int) []int {
	t := len(s)
	for i := 1; i < t; i++ {
		tmp := s[i]
		for j := i - 1; j >= 0; j-- {
			if s[j] <= tmp {
				break
			} else {
				s[j+1] = s[j]
				s[j] = tmp
			}
		}
	}
	return s
}

//选择
//每次选择一个最小值放置，经过N轮循环
//O(N^2)
func Selection(s []int) []int {
	t := len(s)
	for i := 0; i < t-1; i++ {
		min := s[i]
		minIndex := i
		for j := i + 1; j < t; j++ {
			if s[j] < min {
				min = s[j]
				minIndex = j
			}

		}
		if minIndex != i {
			s[i], s[minIndex] = s[minIndex], s[i]
		}
	}
	return s
}

//快排
//O(nlogn)
func QuickSort(s []int) []int {
	return qsort(s)
}

//来自stackoverflow
func qsort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// Pick a pivot
	//	pivotIndex := rand.Int() % len(a)
	pivotIndex := 0

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	//找到left个比right小得数
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	qsort(a[:left])
	qsort(a[left+1:])

	return a
}

//归并
//O(nlogn) 需要额外的存储空间
func MegerSort(a []int) []int {
	return merge(a)
}
func Merge(a, b []int) []int {
	total := len(a) + len(b)
	r := make([]int, total)
	la := len(a)
	lb := len(b)
	i, j, k := 0, 0, 0
	for ; i < la && j < lb; k++ {
		if a[i] <= b[j] {
			r[k] = a[i]
			i++
		} else {
			r[k] = b[j]
			j++
		}
	}
	if i < la {
		copy(r[(total-la+i):], a[i:])
	}
	if j < lb {
		copy(r[(total-lb+j):], b[j:])
	}
	return r
}

func merge(a []int) []int {
	if len(a) == 1 {
		return a
	}
	low, high := 0, len(a)-1
	mid := (low+high)/2 + 1
	a1 := merge(a[low:mid])
	a2 := merge(a[mid : high+1])
	return Merge(a1, a2)
}

//堆排序
//O(nlgn)
func HeapSort(a []int) []int {
	createHeap(a)
	t, index := len(a), 1
	for index < t {
		a[0], a[t-index] = a[t-index], a[0]
		adjustHeap(a, 0, t-index)
		index++
	}
	return a
}

//创建堆
func createHeap(a []int) {
	t, mid := len(a), len(a)/2+1
	for i := mid; i >= 0; i-- {
		adjustHeap(a, i, t)
	}
}

//调整堆
//大顶堆
func adjustHeap(a []int, start, aL int) []int {
	left := 2 * start
	right := 2*start + 1
	if left >= aL {
		return a
	}
	maxIndex := start
	if a[left] > a[maxIndex] {
		maxIndex = left
	}
	if right < aL && a[right] > a[maxIndex] {
		maxIndex = right
	}
	if maxIndex != start {
		a[start], a[maxIndex] = a[maxIndex], a[start]
		adjustHeap(a, maxIndex, aL)
	}
	return a
}

//基数
//分位数比较，从个位开始比较
//获取最大位数
func getMaxCount(a []int) int {
	if len(a) == 0 {
		return 0
	}
	maxIndex := 0
	for k := 1; k < len(a); k++ {
		if a[k] > a[maxIndex] {
			maxIndex = k
		}
	}
	count := 0
	t := a[maxIndex]
	for t != 0 {
		count++
		t = t / 10
	}
	return count
}

//获取相应位置得数字
func getDigitFromIndex(t int, index int) int {
	for index != 1 {
		t = t / 10
		index--
	}
	return t % 10
}

//将整数按照位数分割成不同的数字，然后每个位数分别比较
//会使用到桶，通过将要比较的位，将要排序的元素分配到0-9个桶中
//O(kn) 需要额外的存储空间
func RadixSort(a []int) []int {
	t := len(a)
	maxCount := getMaxCount(a)
	if maxCount == 0 {
		return a
	}
	bucketsMap := make(map[int][]int)
	for i := 0; i < 10; i++ {
		bucketsMap[i] = make([]int, 0)
	}
	for i := 0; i < maxCount; i++ {
		result := make([]int, 0)
		for j := 0; j < t; j++ {
			k := getDigitFromIndex(a[j], i+1)
			bucketsMap[k] = append(bucketsMap[k], a[j])
		}
		for k := 0; k < 10; k++ {
			result = append(result, bucketsMap[k]...)
			bucketsMap[k] = make([]int, 0)
		}
		a = result
	}
	return a
}

//计数
//找出最大值和最小值，根据差值生成数组记录期间有多少个数
//空间和时间复杂度均为O(n+k)
func getMaxMin(a []int) (int, int) {
	if len(a) == 0 {
		return 0, 0
	}
	min, max := a[0], a[0]
	for k, v := range a {
		if v > max {
			max = a[k]
		}
		if v < min {
			min = a[k]
		}
	}
	return max, min
}
func CountSort(a []int) []int {
	t := len(a)
	if t == 0 {
		return a
	}
	//得到最大值和最小值
	min, max := a[0], a[0]
	for k, v := range a {
		if v > max {
			max = a[k]
		}
		if v < min {
			min = a[k]
		}
	}
	count := make(map[int]int)
	for _, v := range a {
		index := v - min
		count[index] = count[index] + 1
	}
	index := 0
	for i := 0; i < (max-min)+1; i++ {
		for k := 0; k < count[i]; k++ {
			a[index] = min + i
		}
		index += count[i]
	}
	return a
}

//桶排序
//计数排序的升级版
//桶中元素需要进行排序
//平均时间复杂度O(n+k) 最坏O(n^2)
func BucketSort(a []int) []int {
	if len(a) == 0 {
		return a
	}
	max, min := getMaxMin(a)
	//桶的初始化
	var defaultBucketSize = 100
	var bucketCount = (max-min)/defaultBucketSize + 1
	buckets := make(map[int][]int)
	for i := 0; i < bucketCount; i++ {
		buckets[i] = make([]int, 0)
	}
	for i := 0; i < len(a); i++ {
		index := (a[i] - min) / defaultBucketSize
		buckets[index] = append(buckets[index], a[i])
	}
	result := make([]int, 0)
	for i := 0; i < bucketCount; i++ {
		fmt.Println(len(buckets[i]))
		Insert(buckets[i])
		result = append(result, buckets[i]...)
	}
	return result
}

//希尔排序
//插入排序的一种更高效的改进版本
//先将整个待排序的记录序列分割成若干子序列分别直接进行插入排序，待整个序列中的记录基本有序时再对全进行直接插入
func ShellSort(a []int) []int {
	if len(a) == 0 {
		return a
	}
	gap := 1
	temp := len(a)
	for gap < len(a)/3 {
		gap = gap*3 + 1
	}
	for ; gap > 0; gap = gap / 3 {
		for i := gap; i < len(a); i++ {
			temp = a[i]
			j := i - gap
			for ; j >= 0 && a[j] > temp; j = j - gap {
				a[j+gap] = a[j]
			}
			a[j+gap] = temp
		}
	}
	return a
}
