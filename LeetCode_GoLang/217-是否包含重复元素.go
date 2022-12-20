package LeetCode_GoLang

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, val := range nums {
		_, ok := m[val]
		if !ok {
			m[val] = 0
		}
		m[val]++
		if m[val] > 1 {
			return true
		}
	}
	return false
}
