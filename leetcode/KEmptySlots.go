package main                                                                                                                                                                                                                                                                  

import "fmt"

// There is a garden with N slots. In each slot, there is a flower. The N flowers will bloom one by one in N days. In each day, there will be exactly one flower blooming and it will be in the status of blooming since then.

// Given an array flowers consists of number from 1 to N. Each number in the array represents the place where the flower will open in that day.

// For example, flowers[i] = x means that the unique flower that blooms at day i will be at position x, where i and x will be in the range from 1 to N.

// Also given an integer k, you need to output in which day there exists two flowers in the status of blooming, and also the number of flowers between them is k and these flowers are not blooming.

// If there isn't such day, output -1.

// Example 1:
// Input: 
// flowers: [1,3,2]
// k: 1
// Output: 2
// Explanation: In the second day, the first and the third flower have become blooming.
// Example 2:
// Input: 
// flowers: [1,2,3]
// k: 1
// Output: -1


//思路：苯方法，利用map记录已经开的花，每前进一天，遍历map查看是否符合要求，第K天，位置N，则查看N-k-1 和 N+k+1 是否已经开（还需要验证位置是否合法),如何有开的，则查询两者之间是否有开放的，如果没有，则满足情况

// 位置，天数，k值，
func main() {
	   fmt.Println("vim-go")
	   flowers := []int{3, 9, 2, 8, 1, 6, 10, 5, 4, 7}
	   fmt.Println(kEmptySlots(flowers, 1))
}
func kEmptySlots(flowers []int, k int) int {
	   records := make(map[int]bool)
	   for i := 0; i < len(flowers); i++ {
			   pos := flowers[i]
			   records[pos] = true
			   if _, ok := records[pos-k-1]; ok {
					   exist := false
					   for ii := pos - k; ii < pos; ii++ {
							   if _, ok := records[ii]; ok {
									   exist = true
									   break
							   }
					   }
					   if !exist {
							   return i + 1                                                                                                                                                                                                                                                          }                                                                                  
			   }                                                                                                                                                                                                                                                                             if _, ok := records[pos+k+1]; ok {                                                                                 
					   exist := false                                                                                                                                                                                                                                                                for ii := 1; ii <= k; ii++ {                                                                                               
							   if _, ok := records[ii+pos]; ok {                                                                                                                                                                                                                                                     exist = true                                                                                                                                                       break                                                                                                                                                                                                        
							   }                                                                                                                                                                                                                                                                     }                                                                                                                                                                          
					   if !exist {                                                                                                                                                                                                                                                                           return i + 1                                                                                                                                                                               
					   }                                                                                                                                                                                                                                                                     }                                                                                                                                                                                                                                                   
	   }                                                                                                                                                                                                                                                                             return -1                                                                                                                                                                                                                                                                                                  
}
