//想了好长时间实在没思路，就采用了最笨的方法，两次遍历，依次求出了每个点最大的面积
func largestRectangleArea(heights []int) int {
	
		   hLen := len(heights)
		   if hLen == 0 { 
				   return 0
		   }   
		   if hLen == 1 { 
				   return heights[0]
		   }   
		   max_sum := heights[0]
		   for i := 0; i < hLen; i++ {
				   distance := 1
				   min := heights[i]
				   curIndex := i
				   if heights[curIndex] > max_sum {
						   max_sum = heights[curIndex]
				   }   
				   sum := heights[curIndex]
				   for j := i + 1; j < hLen; j++ {
						   distance += 1
						   if heights[j] < min {
								   min = heights[j]
						   }   
						   t := distance * min 
						   if t >= sum {
								   sum = t 
						   }   
	
				   }   
				   if sum > max_sum {
						   max_sum = sum 
				   }   
		   }   
		   return max_sum
}                         
//通过后参考别人的思路，遍历，使用栈，这个是把当前的元素作为最小元素，求出以当前元素为最小元素能够得出的最大面积
//遇见比栈顶高度大的，取出来，算面积，然后依次退栈直到没有元素或者栈顶元素高度不大于当前的高度。得出此轮的最大面积。然后把当前元素入栈或者栈顶元素数量加1.
//下轮循环开始
//for循环完毕，得到一个最大的面积。此时还需要计算栈中存储的元素，栈中存储的元素，算是根据高度排列好的，再计算下是否比当前最大元素大。
//比较完后得出最大的元素
type item struct {                
	h     int                 
	count int                 
}                                 
							  
func largestRectangleArea(heights []int) (ans int) {
	var stack []item          
	for _, h := range heights {
			fmt.Println(stack)                                                                                                                                                                                                                                            
			top := len(stack) - 1
			cnt := 0          
			//找出了比栈顶元素更小的元素，这里就需要计算当前栈顶元素能够达到的最大面积
			//栈中要存储比栈顶元素大的数据，知道遇见比栈顶元素小的数据开始计算以当前栈顶元素为最小元素，矩阵面积能够达到的最大值
			for top >= 0 && stack[top].h > h {
					cnt += stack[top].count
					if a := cnt * stack[top].h; a > ans {
							ans = a
					}         
					stack = stack[:top]
					top--     
			}                 
			cnt++             
			if top >= 0 && stack[top].h == h {
					stack[top].count += cnt
			} else {          
					stack = append(stack, item{h: h, count: cnt})
			}                 
	}                         
	cnt := 0                  
	fmt.Println(stack)        
	for i := len(stack) - 1; i >= 0; i-- {
			cnt += stack[i].count
			if a := cnt * stack[i].h; a > ans {
					ans = a   
			}                 
	}                         
	return                    
}                                 
func main() {                     
	fmt.Println(largestRectangleArea([]int{2, 3, 5, 6,1, 2, 3}))
}  