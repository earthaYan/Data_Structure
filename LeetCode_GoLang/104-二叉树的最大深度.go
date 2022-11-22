package LeetCode_GoLang

func computeDepth(root *TreeNode, currentDepth int) int {
	if root == nil {
		return currentDepth
	}
	// 计算左子树的深度
	nLeft := computeDepth(root.Left, currentDepth+1)
	// 计算右子树的深度
	nRight := computeDepth(root.Right, currentDepth+1)
	//获取最大深度
	if nLeft > nRight {
		return nLeft
	} else {
		return nRight
	}
}
func maxDepth(root *TreeNode) int {
	// 二叉树的最大深度
	// 记录当前深度
	return computeDepth(root, 0)

}
