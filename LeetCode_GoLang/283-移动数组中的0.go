package LeetCode_GoLang

func moveZeroes(nums []int) {
	zeroSize := 0
	for index, val := range nums {
		if val == 0 {
			zeroSize++
		} else if zeroSize > 0 {
			temp := nums[index]
			nums[index] = 0
			nums[index-zeroSize] = temp
		}
	}
}
