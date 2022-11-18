package LeetCode_GoLang

func postorderTraversal(root *TreeNode) []int {
	// LRD
	var nodes = make([]int, 0)
	if root == nil {
		return nodes
	}
	if root.Left != nil {
		nodes = append(nodes, inorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		nodes = append(nodes, inorderTraversal(root.Right)...)
	}
	nodes = append(nodes, root.Val)
	return nodes
}
