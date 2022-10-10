package LeetCode_GoLang

import "fmt"

func max(nums []int) int {
	var maxVal int = nums[0]
	for _, val := range nums {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}
func MaxSubArray(nums []int) int {
	// 未通过
	/**
	 	link:https://leetcode.cn/problems/maximum-subarray/
		给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
		子数组 是数组中的一个连续部分。
	*/
	// 1.如果为空数组，则返回0
	// 2.如果只有一个数组元素，返回当前元素
	// 3.如果只有两个数组元素且有负数，则返回最大值
	// 4.最大和的连续子数组即要求子数组中不能超过1个负数，如超过则结束本轮循环，进入下一轮，同时将子数组和放入新的数组中
	// 5.比较所有的子数组和，返回最大值
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 && nums[0] < 0 || nums[1] < 0 {
		return max(nums)
	}
	negative := 0
	sum := 0
	var sumArr []int
	for key, val := range nums {
		if val < 0 {
			negative++
		}
		sumArr = append(sumArr, sum+val)
		if negative > 1 {
			sum = 0
			negative = 0
			continue
		} else if key == len(nums)-1 {
			break
		}
	}
	fmt.Println(sumArr)
	return max(sumArr)
}
