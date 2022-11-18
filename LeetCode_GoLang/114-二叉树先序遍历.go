package LeetCode_GoLang

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var nodes = make([]int, 0)
	//DLR
	if root == nil {
		return nodes
	}
	nodes = append(nodes, root.Val)
	if root.Left != nil {
		nodes = append(nodes, preorderTraversal(root.Left)...)
	}
	if root.Right != nil {
		nodes = append(nodes, preorderTraversal(root.Right)...)
	}
	return nodes
}

func ExecutePreorderTraversal() {
	arr := &TreeNode{
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
	nodes := preorderTraversal(arr)
	fmt.Println(nodes)
}
