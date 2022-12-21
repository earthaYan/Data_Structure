package LeetCode_GoLang

func ThirdMax(nums []int) int {
	//降序排序后去重，返回第三个元素或第一个元素
	arr := removeDuplicateElement(sort1(nums))
	if len(arr) < 3 {
		return arr[0]
	}
	return arr[2]
}

func removeDuplicateElement(nums []int) []int {
	result := make([]int, 0, len(nums))
	temp := map[int]struct{}{}
	for _, item := range nums {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
func sort1(nums []int) []int {
	var x int
	var j int
	for index := 1; index < len(nums); index++ {
		x = nums[index]                                 //空出a[i]的位置
		for j = index - 1; j >= 0 && x > nums[j]; j-- { //移动定位
			nums[j+1] = nums[j]
		}
		nums[j+1] = x //将x插入在j+1位置
	}
	return nums
}
