package LeetCode_GoLang

/**
给定一个包含 n 个整数的数组 nums，其中 nums[i] 在 [1, n] 范围内，
返回一个包含 [1, n] 范围内所有没有出现在 nums 中的整数的数组
*/
func FindDisappearedNumbers(nums []int) []int {
	miss := []int{}
	for index,_ := range nums {
		if searchVal(nums, index+1) == false {
			miss = append(miss, index+1)
		}
	}
	return miss
}
func searchVal(targets []int, val int) bool {
	for _,value:= range targets {
		if value== val {
			return true
		}
	}
	return false
}
