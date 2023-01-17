package LeetCode_GoLang

import (
	"math"
	"strconv"
	"strings"
)

func DifferenceOfSum(nums []int) int {
	elementSum := 0
	digitSum := 0
	newArr := []string{}
	for i := 0; i < len(nums); i++ {
		ele := strconv.Itoa(nums[i])
		digits := strings.Split((ele), "")
		newArr = append(newArr, digits...)
		elementSum = elementSum + nums[i]
	}

	for j := 0; j < len(newArr); j++ {
		res, err := strconv.Atoi(newArr[j])
		if err == nil {
			digitSum = digitSum + res
		}
	}

	return elementSum - digitSum
}

// 题解上其他解法：
func differenceOfSum(nums []int) int {
	s1, s2 := 0, 0
	for _, num := range nums {
		s1 += num
		tmp := num
		for tmp > 0 {
			s2 += tmp % 10
			tmp /= 10
		}
	}

	return int(math.Abs(float64(s1 - s2)))
}
