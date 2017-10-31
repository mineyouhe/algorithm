//就是依次判断在本次能跳转的最大范围内，哪个位置能够往后跳转的多，然后再以这个位置向后跳转
func jump(nums []int) int {
	if len(nums) == 1 { 
			return 0
	}   
		
	return next(nums, 0)
}           
		
func next(nums []int, index int) int {
	if (nums[index] + index) >= (len(nums) - 1) {                                                                                                       
			return 1
	} else {
			max := nums[index+1]
			max_index := index + 1 
			for i := index + 2; i <= (nums[index] + index); i++ {
					if (max + max_index) <= (nums[i] + i){
							max = nums[i]
							max_index = i 
					}   
		
			} 
			return next(nums, max_index) + 1 
	}   
} 