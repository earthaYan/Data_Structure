package course

func MergeArr(arr1 []int, arr2 []int) []int {
	up, down := 0, 0
	var merged []int
	for up < len(arr1) && down < len(arr2) {
		if arr1[up] < arr2[down] {
			merged = append(merged, arr1[up])
			up++
		} else if arr1[up] > arr2[down] {
			merged = append(merged, arr2[down])
			down++
		} else {
			merged = append(merged, arr2[down])
			down++
			up++
		}
	}
	// 复制后半段
	if up<len(arr1){
		merged=append(merged, arr1[up:]...)
	}else if down<len(arr2){
		merged=append(merged, arr2[down:]...)
	}
	return merged
}
