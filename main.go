package main

import (
	"fmt"
	"leet_code_go/LeetCode_GoLang"
)

func main() {

}
func init() {
	nums1:= []int{1,2}
	nums2:=[]int{3,4}
	res := LeetCode_GoLang.FindMedianSortedArrays(nums1,nums2)
	fmt.Println(res)
}
