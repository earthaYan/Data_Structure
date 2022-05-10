const Status = {
  OK: 1,
  ERROR: 0,
  TRUE: 1,
  FALSE: 0,
};
class Node {
  constructor(data) {
    // 数据域：存放数据元素
    this.data = data;
    // 指针域：存放直接后继节点的地址
    this.next = null;
  }
}
class ChainList {
  constructor() {
    this.size = 0;
    // 记录单链表的长度或节点个数
    this.head = new Node('head');
    this.currentNode = '';
  }
  getElem(index) {
    // 声明一个指针P,指向链表的第一个节点,初始化j为1
    this.currentNode = this.head;
    let j = 1;
    // 当j<i的时候,遍历链表,让p的指针向后移动,不断指向下一个节点,j累加
    while (this.currentNode && j < index) {
      this.currentNode = this.currentNode.next;
      ++j;
    }
    // 如果到链表末尾,p为空,则说明第i个节点不存在
    if (!this.currentNode || j > index) {
      return Status.ERROR;
    }
    // 否则查找成功,返回节点p的数据
    return `${Status.OK}:${JSON.stringify(this.currentNode)}`;
  }
  insertElem(index, elem) {
    // 声明指针p指向链表头节点,初始化j从1开始
    this.currentNode = this.head;
    let j = 1;
    // 当j<i时,遍历列表,让p的指针向后移动,不断指向下一节点,j累加
    while (this.currentNode && j < index) {
      this.currentNode = this.currentNode.next;
      j++;
    }
    // 如果到链表末尾p为空,则说明第i个节点不存在
    if (!this.currentNode || j > index) {
      return Status.ERROR;
    }
    // 否则查找成功,在系统中生成空节点s, 将数据元素e赋值给s->data
    let newNode = new Node(elem);
    // 单链表的插入标准语句
    newNode.next = this.currentNode.next;
    this.currentNode.next = newNode;
    this.size++;
    // 返回成功
    return Status.OK;
  }
  delElem(index) {
    // 声明指针p指向链表头节点,初始化j从1开始
    this.currentNode = this.head;
    let j = 1;
    // 当j<i时,遍历列表,让p的指针向后移动,不断指向下一节点,j累加
    while (this.currentNode && j < index) {
      this.currentNode = this.currentNode.next;
      ++j;
    }
    // 如果到链表末尾p为空,则说明第i个节点不存在
    if (!this.currentNode || j > index) {
      return Status.ERROR;
    }
    // 否则查找成功,将需要删除的节点p->next赋值给q
    let delNode = this.currentNode.next;
    // 单链表的删除标准语句
    this.currentNode.next = delNode.next;
    // 将q节点的数据赋值给e,作为返回
    const data = delNode.data;
    // 释放q节点
    delNode = null;
    this.size--;
    // 返回成功
    return `${Status.OK} data`;
  }
  display() {
    let result = '';
    let currNode = this.head;
    while (currNode) {
      result += currNode.data;
      currNode = currNode.next;
      if (currNode) {
        result += '->';
      }
    }
    console.log(result);
  }
  createListToHead(n) {
    // 单链表的整表创建
    // 声明指针p和计数器遍历i
    // 初始化一空链表L
    // 让L的头节点的指针指向null,即建立一个带头节点的单链表
    // 循环
    for (let i = 0; i < n; i++) {
      // 1.生成新节点赋值给p
      // 2.随机生成一个数字赋值给p的数据域p->data
      const data = Math.floor(Math.random() * 100 + 1);
      console.log(data);
      let newNode = new Node(data);
      // 3.将p插入到头节点与前一新节点之间
      newNode.next = this.head.next;
      this.head.next = newNode;
      this.size++;
    }
  }
  createListToTail(n) {
    this.currentNode = this.head;
    for (let i = 0; i < n; i++) {
      const data = Math.floor(Math.random() * 100 + 1);
      console.log(data);
      let newNode = new Node(data);
      this.currentNode.next = newNode;
      this.currentNode = this.currentNode.next;
      this.size++;
    }
  }
}
let myList = new ChainList();
myList.createListToTail(3);
myList.display();
console.log(JSON.stringify(myList));
