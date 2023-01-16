package course


func SequentialSearchFromLeft(arr []int, ele int) int {
	// 顺序查找
	// 以表的左侧为起点，向另一个端点逐个元素查找
	for left := 0; left < len(arr); left++ {
		if arr[left] == ele {
			return left
		}
	}
	return -1
}
func SequentialSearchFromRight(arr []int, ele int) int {
	// 顺序查找
	// 以表的右端为起点，向另一个端点逐个元素查找
	for right := len(arr) - 1; right >= 0; right-- {
		if arr[right] == ele {
			return right
		}
	}
	return -1
}

// 优化方法：在查找终点预设监督元节点
func SequentialSearchFromLeftUsingNode(arr []int, ele int) int {
	// 顺序查找-使用监督元节点
	// 以表的左侧为起点，向另一个端点逐个元素查找
	newArr := append(arr[:], ele)
	left := 0
	for newArr[left] != ele {
		left++
	}
	if left < len(arr) {
		return left
	}
	return -1
}
func SequentialSearchFromRightUsingNode(arr []int, ele int) int {
	// 顺序查找-使用监督元节点
	newArr := []int{ele}
	newArr=append(newArr, arr...)
	right:=len(newArr)-1
	for newArr[right]!=ele{
		right--
	}
	if right>0{
		return right-1
	}
	return -1
}
