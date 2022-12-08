package LeetCode_GoLang

import (
	"fmt"
)

func RemoveElement(nums []int, val int) int {
	// 0, 1, 2, 2, 3, 0, 4, 2|||2
	// 0,1,2,3,0,2,4,2,_
	// 0,1,3,0,4,2,_,_
	// 0,1,3,0,4,_,_,_
	// 出错点：无法处理连续重复的元素
	// 解决方法：暂时使用goto语句
	count := len(nums)
	//删除指定值，返回剩下的元素个数
LOOP:
	for index, value := range nums {
		if value == val {
			// 后面所有元素前移1位，最后一位元素置为空
			for i := index; i < len(nums)-1; i++ {
				nums[i] = nums[i+1]
			}
			nums[len(nums)-1] = '_'
			count--
			goto LOOP
		} else {
			continue
		}
	}
	fmt.Print(nums)
	return count
}
