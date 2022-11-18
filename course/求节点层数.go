package course

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preOrder(root *TreeNode, search *TreeNode, currentDepth int) int {
	// 递归遍历的时候传递一个表示当前上一层节点深度的参数
	if root == nil {
		return 0
	}
	currentDepth++
	if root == search {
		// 当找到目标节点的时候，返回当前深度
		return currentDepth
	}
	if root.Left != nil {
		preOrder(root.Left, search, currentDepth)
	}
	if root.Right != nil {
		preOrder(root.Right, search, currentDepth)
	}
	return currentDepth + 1
}

// 求指定节点的层数可以转换为求指定节点的深度+1，即从root节点到指定节点的边数
func getLayer(rootNode *TreeNode, searchNode *TreeNode) int {
	if searchNode == nil {
		return 0
	}
	depth := preOrder(rootNode, searchNode, 0)
	return depth
}

func ExecuteGetLayer() {
	arr2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 9,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	nodes := getLayer(arr2, arr2.Left.Left)
	fmt.Print(nodes)
}
