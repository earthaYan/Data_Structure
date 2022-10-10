package LeetCode_GoLang

func TwoSum(nums []int, target int) []int {
	/**
	 	link:https://leetcode.cn/problems/two-sum
		给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
	*/
	// 1.设置一个空数组
	// 2.双重循环，将第一层当前数字与后面的数字依次相加，如果等于target，则将下标插入数组中，跳出循环
	targetPos := []int{}
ref:
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				targetPos = []int{i, j}
				break ref
			}
		}
	}
	return targetPos
}
