package LeetCode_GoLang

import "sort"

func SortPeople(names []string, heights []int) []string {
	m := make(map[int]string, len(names))
	// map 关联用户名和身高
	for index, val := range heights {
		m[val] = names[index]
	}
	keys := make([]int, 0, len(heights))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	res := make([]string, len(names))
	j := len(names) - 1

	for _, k := range keys {
		res[j] = m[k]
		j--
	}
	return res
}
