package course

import (
	"math"
)

// 二分查找思想：缩减问题规模
func BinarySearch(arr []int, ele, left, right int) int {
	if right < left {
		return -1
	}
	// 计算中值位置
	mid := int(math.Floor(float64((left + right) / 2)))
	if ele == arr[mid] {
		return mid
	} else if ele < arr[mid] {
		// 舍弃右侧
		return BinarySearch(arr, ele, left, mid-1)
	} else {
		// 舍弃左侧
		return BinarySearch(arr, ele, mid+1, right)
	}
}
