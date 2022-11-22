package LeetCode_GoLang

func swap(leftTree *TreeNode, rightTree *TreeNode, data int) *TreeNode {
	// 构建一个新节点，其左右子树分别为原来的子树的相反方向
	p := &TreeNode{Val: data}
	p.Right = leftTree
	p.Left = rightTree
	return p
}
func invertTree(root *TreeNode) *TreeNode {
	// 本质上貌似还是后续遍历?
	if root == nil {
		return root
	}
	nLeft := invertTree(root.Left)
	nRight := invertTree(root.Right)
	return swap(nLeft, nRight, root.Val)
}
