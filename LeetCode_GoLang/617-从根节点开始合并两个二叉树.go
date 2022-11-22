package LeetCode_GoLang

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	// 递归终止条件
	if root1 == nil {
		return root2
	} else if root2 == nil {
		return root1
	}
	root1.Val += root2.Val
	root1.Left = mergeTrees(root1.Left, root2.Left)
	root1.Right = mergeTrees(root1.Right, root2.Right)
	return root1
}
