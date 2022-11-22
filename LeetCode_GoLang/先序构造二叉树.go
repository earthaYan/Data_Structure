package LeetCode_GoLang

import (
	"math/rand"
	"time"
)

func insert(t *TreeNode, v int) *TreeNode {
	// 检查操作树是否为空
	if t == nil {
		return &TreeNode{v, nil, nil}
	}
	// 判断二叉树上是否已经存在将要插入的值。如果值已经存在，那么函数将什么也不做然后返回
	if v == t.Val {
		return t
	}
	if v < t.Val {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}
func createTreePre(n int) *TreeNode {
	var t *TreeNode
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*n; i++ {
		temp := rand.Intn(n * 2)
		t = insert(t, temp)
	}
	return t
}
