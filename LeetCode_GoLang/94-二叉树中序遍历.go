package LeetCode_GoLang

import "fmt"

func inorderTraversal(root *TreeNode) []int {
	nodes := make([]int, 0)
	// LDR
	if root == nil {
		return nodes
	}
	if root.Left != nil {
		nodes = append(nodes, inorderTraversal(root.Left)...)
	}
	nodes = append(nodes, root.Val)
	if root.Right != nil {
		nodes = append(nodes, inorderTraversal(root.Right)...)
	}
	return nodes
}
func ExecuteInorderTraversal() {
	arr1 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
		},
	}

	nodes := inorderTraversal(arr1)
	fmt.Println(nodes)
}
