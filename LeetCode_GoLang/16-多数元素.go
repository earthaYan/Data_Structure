package LeetCode_GoLang

func majorityElement(nums []int) int {
	m := make(map[int]int)
	maxKey := nums[0]
	for _, val := range nums {
		_, ok := m[val]
		if !ok {
			m[val] = 0
		}
		m[val]++
	}
	for key, val := range m {
		if val > m[maxKey] {
			maxKey = key
		}
	}
	return maxKey
}
