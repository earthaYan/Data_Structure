package LeetCode_GoLang

// 最小深度指的是：从根节点到叶子节点的最短路径上的节点数量。
// 错误解法：根据最大深度直接设置为取最小值
func computeMinDepthError(root *TreeNode, currentDepth int) int {
	if root == nil {
		return currentDepth
	}
	// 计算左子树的深度
	nLeft := computeMinDepthError(root.Left, currentDepth+1)
	// 计算右子树的深度
	nRight := computeMinDepthError(root.Right, currentDepth+1)
	//获取最大深度
	if nLeft < nRight {
		return nLeft
	} else {
		return nRight
	}
}
func minDepthError(root *TreeNode) int {
	return computeMinDepthError(root, 0)
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 递归左子树的最小高度，递归右子树的最小高度，求根的最小高度
	nLeft := minDepth(root.Left)
	nRight := minDepth(root.Right)
	// 当其中一个子树为空的时候，并不是最低点
	if root.Left == nil {
		return nRight + 1
	} else if root.Right == nil {
		return nLeft + 1
	}
	if nLeft > nRight {
		return nRight + 1
	} else {
		return nLeft + 1
	}
}

// 	二叉树的深度：根节点到最远叶子节点的最长路径上的节点数。
//	最大深度：根节点到最远叶子节点的最长路径上的节点总数。
//  最小深度：根节点到最近叶子节点的最短路径上的节点总数
// 	二者差别主要在于处理左右孩子不为空的逻辑
