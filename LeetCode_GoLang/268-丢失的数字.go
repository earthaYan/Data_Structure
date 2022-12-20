package LeetCode_GoLang

func missingNumber(nums []int) int {
	// 只丢失一个数字，所以可以计算等差数列1+2+...n的值
	// 两个和相减即漏掉的数字
	sum := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		sum += nums[i]
	}
	normalVal := (length + 1) * length / 2
	return normalVal - sum
}
