package course

func BubbleSort(nums []int, n int) []int {
	//反复扫描待排序列，若相邻元素逆序，则交换。直至无逆序
	flag := 1
	j := 0
	for flag > 0 { //存在逆序，进入循环
		flag = 0
		for index := n - 2; index >= j; index-- {
			if nums[index] > nums[index+1] { //发现逆序就交换
				x := nums[index]
				nums[index] = nums[index+1]
				nums[index+1] = x
				flag = 1 //发现交换说明无序
			}
		}
		j++
	}
	return nums
}
