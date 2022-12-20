package LeetCode_GoLang

import "math"

func containsNearbyDuplicate(nums []int, k int) bool {
	// 初始思路：双重循环，判断条件-错误，会超时,所以使用map解决
	m := make(map[int]int, len(nums))
	for index, val := range nums {
		j, ok := m[val]
		if ok && math.Abs(float64(index-j)) <= float64(k) {
			return true
		}
		m[val] = index
	}
	return false
}
