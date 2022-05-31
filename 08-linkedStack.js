Status = {
  OK: 1,
  ERROR: 0,
  TRUE: 1,
  FALSE: 0,
};
class StackNode {
  constructor(data) {
    this.data = data;
    this.next = null;
  }
}
class LinkedStack {
  constructor() {
    this.top = null;
    this.count = 0;
  }
  pushToLinkedStack(data) {
    // 新建节点
    const newNode = new StackNode(data);
    // 把当前栈顶元素的赋值给新节点的直接后继
    newNode.next = this.top;
    // 把新节点设置为栈顶指针
    this.top = newNode;
    this.count++;
    return Status.OK;
  }
  stackEmpty() {
    return !this.count;
  }
  popFromLinkedStack() {
    if (this.stackEmpty()) {
      return Status.ERROR;
    }
    const delData = this.top.data;
    const delNode = this.top;
    // 栈顶指针后移一位
    this.top = delNode.next;
    this.count--;
    return delData;
  }
  display() {
    let result = '';
    let currNode = this.top;
    while (currNode) {
      result += currNode.data;
      currNode = currNode.next;
      if (currNode) {
        result += '->';
      }
    }
    console.log(result);
  }
}
let newStack = new LinkedStack();
newStack.pushToLinkedStack('a');
newStack.pushToLinkedStack('b');
newStack.pushToLinkedStack('c');
newStack.display();
newStack.popFromLinkedStack();
newStack.display();
newStack.popFromLinkedStack();
newStack.display();
