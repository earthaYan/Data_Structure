package LeetCode_GoLang
func ContainsDuplicate(nums []int) bool{
	/**
	 	link:https://leetcode.cn/problems/contains-duplicate
		给你一个整数数组 nums 。
		如果任一值在数组中出现 至少两次 ，返回 true ；
		如果数组中每个元素互不相同，返回 false 。
	*/
	// 1. 设置一个map集合,并循环遍历数组
	// 2. 如果当前元素在map的集合中，则说明存在重复元素，返回true
	// 3. 如果当前元素不在map集合中，则设置map[key]=1
	// 4. 如果map集合中所有的值都为1，则说明不存在重复元素，返回false
	numMap:=make(map[int]int)
	for _,val:=range nums{
		_,isExistInMap:= numMap[val]
		if(isExistInMap){
			return true
		}
		numMap[val]=1
	}
	return false
}