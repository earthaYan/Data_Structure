package LeetCode_GoLang

func sum(root *TreeNode) int {
	currentSum := 0
	if root == nil {
		return currentSum
	}
	currentSum += root.Val
	return currentSum
}
func checkTree(root *TreeNode) bool {
	// 计算左子树之和
	nLeft := sum(root.Left)
	// 计算右子树之和
	nRight := sum(root.Right)
	// 判断两边之和是否等于根节点的值
	return root.Val == nLeft+nRight
}
