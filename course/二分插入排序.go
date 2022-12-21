package course

import "math"

func BinaryInsert(nums []int, n int) []int {
	for index := 1; index < n; index++ {
		x := nums[index] //准备插入元素nums[index]
		left := 0
		right := index - 1  //设置查找段的起点和终点
		for left <= right { //二分定位
			midVal := (left + right) / 2
			mid := int(math.Floor(float64(midVal)))
			if x < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		for j := index - 1; j >= left; j-- {
			nums[j+1] = nums[j] //元素右移
		}
		nums[left] = x //待插入元素就位
	}
	return nums
}
