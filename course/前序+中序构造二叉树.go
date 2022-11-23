package course

func buildTree(pre []int, in []int, pre1 int, pre2 int, in1 int, in2 int) *TreeNode {
	var rootPos int    //指明中序序列中根的位置
	var tree *TreeNode //用于新建节点
	if pre1 > pre2 {
		return nil
	}
	//构造根节点
	tree.Val = pre[pre1]
	//	在中序序列中找根
	for rootPos := in1; rootPos < in2; rootPos++ {
		if pre[rootPos] == in[in1] {
			break
		}
	}
	//	未找到根退出
	//	递归构造左子树
	tree.Left = buildTree(pre, in, pre1+1, pre1+rootPos-in1, in1, rootPos-1)
	// 	递归构造右子树
	tree.Right = buildTree(pre, in, pre1+rootPos-in1+1, pre2, rootPos+1, in2)
	//返回根节点
	return tree
}
func Generate(pre []int, in []int) *TreeNode {
	return buildTree(pre, in, 0, len(pre)-1, 0, len(pre)-1)
}
