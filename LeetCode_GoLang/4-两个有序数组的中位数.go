package LeetCode_GoLang

// 思路：可以先合并两个有序数组，然后求中位数
func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	up, down := 0, 0
	var merged []int
	for up < len(nums1) && down < len(nums2) {
		switch {
		case nums1[up] < nums2[down]:
			merged = append(merged, nums1[up])
			up++
		case nums1[up] > nums2[down]:
			merged = append(merged, nums2[down])
			down++
		default:
			merged = append(merged, nums1[up], nums2[down])
			down++
			up++
		}
	}
	if up < len(nums1) {
		merged = append(merged, nums1[up:]...)
	} else if down < len(nums2) {
		merged = append(merged, nums2[down:]...)
	}
	left := 0
	right := len(merged) - 1
	midPos := (left + right) / 2
	if len(merged)%2 != 0 {
		// 奇数
		return float64(merged[midPos])
	} else {
		// 偶数
		return float64(merged[midPos]+merged[midPos+1]) / 2
	}
}
