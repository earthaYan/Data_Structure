package LeetCode_GoLang

func UnequalTriplets(nums []int) int {
	// 需求：找到满足条件的元组个数
	// 要求：0<=i<j<k<nums.length && 两两互不相等
	// 思路->暴力破解：时间复杂度O(N^3)
	var count int = 0
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			if nums[i] == nums[j] {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				if nums[i] != nums[j] && nums[i] != nums[k] && nums[j] != nums[k] {
					count++
				}
			}
		}
	}
	return count
}

// 题解中大部分都是暴力破解，但票数最高的是时间复杂度为O(N)的解法，待学习
