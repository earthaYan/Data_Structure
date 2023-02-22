package LeetCode_GoLang
/**
给定一个大整数，表示为整数数组 digits，其中每个 digits[i] 是整数的第 i 个数字。 
这些数字按从左到右的顺序从最高有效位到最低有效位排序。 
大整数不包含任何前导 0。 将大整数递增 1 并返回生成的数字数组。
**/
func PlusOne(digits []int) []int {
	nums := make([]int,len(digits))
	copy(nums, digits)
	len := len(nums)
	val:=nums[len-1] + 1
	if val==10{
		nums =append(nums, 1,0)
	}else{
		nums[len-1] =val
	}
	return nums
}