package course

type linkNode struct {
	data int
	next *linkNode
}

// 链表指定位置插入本质：局部修改节点的链域
const End_Elem = 10

// 向前插入法
func NewLinkWithHead() *linkNode {
	var head, newNode *linkNode
	var x int

	newNode = &linkNode{}
	for x != End_Elem {
		newNode = &linkNode{} //申请一个新节点
		newNode.data = x      //将读取的元素值存入新节点的值域
		newNode.next = head   //将新节点插在表头处
		head = newNode
	}
	return head
}

// 向后插入法
func newLinkWithTail() *linkNode {
	var head, newNode, last *linkNode
	var x int
	head, last = &linkNode{}, &linkNode{}
	last.next=nil
	for x != End_Elem {
		newNode = &linkNode{}
		newNode.data = x
		last.next=newNode//插在表尾
		newNode.next=nil//尾节点置空
		last=newNode//修改尾指针
	}
	// 删除辅助头节点
	newNode=head
	head=head.next
	newNode=nil
	return head
}
