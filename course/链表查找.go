package course

// 单向加尾链表的查找
func SearchLinkedList(p, last *linkNode, x int) *linkNode {
	last.data = x //设置监督元
	for p.data != x {
		p = p.next //没找到，继续找
	}
	if p != last {
		return p //不是监督元节点，查找成功
	}
	return nil
}