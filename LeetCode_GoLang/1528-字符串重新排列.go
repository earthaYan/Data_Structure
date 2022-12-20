package LeetCode_GoLang

import "fmt"

func RestoreString(s string, indices []int) string {
	// 错误思路：使用map后排序
	// 原因：map是无序集合
	m := make([]byte, len(indices))
	for index, val := range indices {
		res := s[index]
		m[val] = res
	}
	fmt.Print(string(m))
	return string(m)
}
