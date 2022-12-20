package LeetCode_GoLang

func FindLucky(arr []int) int {
	m := make(map[int]int)
	res := []int{}
	for _, val := range arr {
		if _, found := m[val]; !found {
			m[val] = 1
		} else {
			m[val]++
		}
	}
	for key, val := range m {
		if key == val {
			res = append(res, val)
		}
	}
	if len(res) > 0 {
		maxVal := res[0]
		for _, val := range res {
			if val > maxVal {
				maxVal = val
			}
		}
		return maxVal
	}
	return -1
}
