package course

import "fmt"

type TreeNodeWithHeight struct {
	Val    int
	Height int
	Left   *TreeNodeWithHeight
	Right  *TreeNodeWithHeight
}

func getHeight(root *TreeNodeWithHeight) int {
	if root == nil {
		return 0
	}
	hLeft := getHeight(root.Left)
	hRight := getHeight(root.Right)
	if hLeft > hRight {
		// Leetcode中以节点为1度，所以需要加1
		root.Height = hLeft + 1
	} else {
		root.Height = hRight + 1
	}
	return root.Height
}

func ExecuteGetHeight() {
	arr3 := &TreeNodeWithHeight{
		Val: 2,
		Left: &TreeNodeWithHeight{
			Val: 1,
			Left: &TreeNodeWithHeight{
				Val: 9,
			},
		},
		Right: &TreeNodeWithHeight{
			Val: 3,
		},
	}
	nodes := getHeight(arr3)
	fmt.Print(nodes)
}
