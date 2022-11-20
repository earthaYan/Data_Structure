package LeetCode_GoLang

import "fmt"

func preOrderDepth(root *TreeNode, currentDepth int) int {
	if root == nil {
		return currentDepth
	}
	lefy
}
func maxDepth(root *TreeNode) int {
	depth := 0
	// 节点的深度指的是根节点到指定节点的边数，最大深度即根节点到左右子树叶子节点的边数中的最大值
	if root == nil {
		return 0
	}
	depth = preOrderDepth(root, 1)
	return depth
}
func ExecuteMaxDepth() {
	arr4 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	var realDepth int = maxDepth(arr4)
	fmt.Print(realDepth)
}
