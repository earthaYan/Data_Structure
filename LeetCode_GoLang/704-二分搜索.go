package LeetCode_GoLang

import (
	"math"
)

func BinarySearch(nums []int, target int) int {
	return searchDiGui(nums, 0, len(nums)-1, target)
}

// 递归
func searchDiGui(nums []int, left int, right int, target int) int {
	if left > right {
		return -1
	}
	mid := int(math.Floor(float64((left + right) / 2)))
	if nums[mid] == target {
		return mid
	}
	if target < nums[mid] {
		return searchDiGui(nums, left, mid-1, target)
	} else {
		return searchDiGui(nums, mid+1, right, target)
	}
}

// 非递归解法
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := int(math.Floor(float64((left + right) / 2)))
		if nums[mid] == target {
			return mid
		}
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
