package LeetCode_GoLang

func singleNumber(nums []int) int {
	// 通过map【数字】:出现次数解决
	m := make(map[int]int)
	for _, val := range nums {
		_, ok := m[val]
		if !ok {
			m[val] = 0
		}
		m[val]++
	}
	for key, val := range m {
		if val == 1 {
			return key
		}
	}
	return -1
}
