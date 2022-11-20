package course

import "fmt"

var found int = 0

func descent(root *TreeNode, target *TreeNode) []int {
	var nodes = make([]int, 0)
	// 表示当前的节点是否属于节点a的子树中
	if found == 0 {
		return nodes
	}

	if root.Val == target.Val {
		found = 1 //表示下面以a为根的子树的遍历
	}

	if found == 1 {
		// 添加此节点

	}

	if root.Val == target.Val {
		found = 0
	}
	return nodes
}

func ExecuteDescent() {
	arr3 := &TreeNode{
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

	nodes := descent(arr3, arr3.Left)
	fmt.Print(nodes)
}
