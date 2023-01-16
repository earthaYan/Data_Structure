package course

import "fmt"

// 已知数组a中存放的元素为{2,4,6,8,10,12,14,16,18,20}，用二分查找法查找元素x出现的位置及查询次数。
// 如果元素x在表中，则输出try 查找次数,pos x所在位置（下标）。
// 如果元素x不在表中，则输出try 查找次数,not found。try和pos后均有一空格。
var try = 1

func Binary_Search(arr []int, left, right, x int) string {
	try++
	if right < left {
		return "try " + fmt.Sprintf("%d", try) + "不存在于表中"
	}
	mid := (left + right) / 2
	if x == arr[mid] {
		return "try " + fmt.Sprintf("%d", try) + " pos " + fmt.Sprintf("%d", mid)
	} else if x < arr[mid] {
		try++
		return Binary_Search(arr, left, mid-1, x)
	} else {
		try++
		return Binary_Search(arr, mid+1, right, x)
	}
}
