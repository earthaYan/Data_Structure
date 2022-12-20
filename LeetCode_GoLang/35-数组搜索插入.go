package LeetCode_GoLang

import "math"

func SearchInsert(nums []int, target int) int {
	// 可以用二分搜索
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := int(math.Floor(float64((left + right) / 2)))
		if target == nums[mid] {
			return mid
		}
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
