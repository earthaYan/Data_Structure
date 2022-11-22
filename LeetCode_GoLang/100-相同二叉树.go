package LeetCode_GoLang

type TreeNodeWithSaneTree struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// 前序遍历
	// 先比较根节点是否相同
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	} else if p.Val != q.Val {
		return false
	}
	// 递归比较左右子树
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
