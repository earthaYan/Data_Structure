package course

func partition(list []int, left, right int) int {
	pivot := list[left] //导致 left 位置值为空
	for left < right {
		//right指针值 >= pivot right指针左移
		for left < right && pivot <= list[right] {
			right--
		}
		//填补left位置空值
		//right指针值 < pivot right值 移到left位置
		//right 位置值空
		list[left] = list[right]
		//left指针值 <= pivot left指针右移
		for left < right && pivot >= list[left] {
			left++
		}
		//填补right位置空值
		//left指针值 > pivot left值 移到right位置
		//left位置值空
		list[right] = list[left]
	}
	//pivot 填补 left位置的空值
	list[left] = pivot
	return left
}

func QuickSort(list []int, left, right int) []int {
	if right > left {
		//位置划分
		pivot := partition(list, left, right)
		//左边部分排序
		QuickSort(list, left, pivot-1)
		//右边排序
		QuickSort(list, pivot+1, right)
	}
	return list
}
