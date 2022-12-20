package LeetCode_GoLang

import (
	"fmt"
)

func MergeSortArr(nums1 []int, m int, nums2 []int, n int) {
	res := make([]int, 0, m+n)
	i, j := 0, 0
	for {
		if i == m {
			res = append(res, nums2[j:]...)
			break
		}
		if j == n {
			res = append(res, nums1[i:]...)
			break
		}
		if nums1[i] < nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}
	copy(nums1, res)
	fmt.Print(nums1)
}
