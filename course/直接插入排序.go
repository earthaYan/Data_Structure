package course

func InsertSort(arr []int, n int) []int {
	//顺序查找
	// 在有序序列中插入一个元素，保持序列有序
	// 插入前：0~i-1是有序段，i~n-1是无序段
	var x int
	var j int
	for index := 1; index < n; index++ {
		x = arr[index]                                 //空出a[i]的位置
		for j = index - 1; j >= 0 && x < arr[j]; j-- { //移动定位
			arr[j+1] = arr[j]
		}
		arr[j+1] = x //将x插入在j+1位置
	}
	return arr
}
