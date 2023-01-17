package LeetCode_GoLang

func TwoSum1(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}
	return []int{}
}

// 其他解法：无尽循环,使用左右指针
func twoSum(nums []int, t int) []int {
	for l, r := 0, len(nums)-1; ; {
		switch {
		case nums[l]+nums[r] > t:
			r--
		case nums[l]+nums[r] < t:
			l++
		default:
			return []int{l + 1, r + 1}
		}
	}
}
