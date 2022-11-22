package LeetCode_GoLang

func isMirror(leftTree *TreeNode, rightTree *TreeNode) bool {
	if leftTree == nil && rightTree == nil {
		return true
	} else if leftTree == nil || rightTree == nil {
		return false
	} else {
		if leftTree.Val == rightTree.Val {
			return isMirror(leftTree.Left, rightTree.Right) && isMirror(leftTree.Right, rightTree.Left)
		} else {
			return false
		}
	}
}
func isSymmetric(root *TreeNode) bool {
	// 返回true的条件：左子树的右子树==右子树的左子树&&左子树的左子树==右子树的右子树||空树
	tLeft := root.Left
	tRight := root.Right
	// 递归遍历左右子树
	return isMirror(tLeft, tRight)
}
