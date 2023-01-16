package course

import "fmt"

// 已知有两个等长的非降序序列S1, S2, 设计函数求S1与S2并集的中位数。有序序列A[0]，A[1]，...，A[N-1]  的中位数指A[(N-1)/2]
func MidPosNum(s1, s2 []int) int {
	// 合并得到新数组
	up, down := 0, 0
	s1Len, s2Len := len(s1), len(s2)
	var merged []int
	for up < s1Len && down < s2Len {
		if s1[up] < s2[down] {
			merged = append(merged, s1[up])
			up++
		} else if s1[up] > s2[down] {
			merged = append(merged, s2[down])
			down++
		} else {
			merged = append(merged, s2[down])
			up++
			down++
		}
	}
	if up < s1Len {
		merged = append(merged, s1[up:]...)
	} else if down < s2Len {
		merged = append(merged, s2[down:]...)
	}
	
	// 求中位数
	res := merged[len(merged)/2]
	fmt.Println(merged,res)
	return res
}
