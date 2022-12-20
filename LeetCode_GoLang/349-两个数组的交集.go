package LeetCode_GoLang

func intersection(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int)
	// m2 := make(map[int]int)
	res := []int{}
	for _, val := range nums1 {
		if _, found := m1[val]; !found {
			m1[val] = 1
		}
	}
	// for _, val := range nums2 {
	// 	if _, found := m2[val]; !found {
	// 		m2[val] = 1
	// 	}
	// }
	// for key := range m1 {
	// 	if _, found_m2 := m2[key]; !!found_m2 {
	// 		res = append(res, key)
	// 	}
	// }
	// 根据题解优化解法
	for _, val := range nums2 {
		if m1[val] == 1 {
			res = append(res, val)
			m1[val] = 0
		}
	}
	return res
}
