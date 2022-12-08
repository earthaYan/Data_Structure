package LeetCode_GoLang

func RemoveDuplicates(nums []int) int {
	// [1,1,1,2,3]->1,2,3
	// 遍历数组
	i := 1
	for j := 1; j < len(nums); j++ {
		if nums[j-1] != nums[j] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}
