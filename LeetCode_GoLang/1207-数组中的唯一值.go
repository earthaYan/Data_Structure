package LeetCode_GoLang

import "fmt"

/**
给定一个整数数组 arr，如果数组中每个值的出现次数是唯一的，则返回 true，否则返回 false
**/

func UniqueOccurrences(arr []int) bool {
	mapNum := make(map[int]int, len(arr))
	for _, val := range arr {
		_, ok := mapNum[val]
		if ok {
			mapNum[val] = mapNum[val] + 1
		} else {
			mapNum[val] = 1
		}
	}
	fmt.Print(mapNum)
	nums := []int{}
	for _, val := range mapNum {
		nums = append(nums, val)
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return false
			}
		}
	}
	return true
}
