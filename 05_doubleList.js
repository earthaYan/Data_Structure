const Status = {
  OK: 1,
  ERROR: 0,
  TRUE: 1,
  FALSE: 0,
};
class Node {
  constructor(data) {
    this.data = data;
    this.next = null;
    this.prev = null;
  }
}
class doubleList {
  constructor() {
    this.head = null;
    this.tail = null;
    this.length = 0;
  }
  insertElem(position, data) {
    if (position < 0 || position > this.length) return Status.ERROR;
    let newNode = new Node(data);
    // 原链表为空
    // 情况1：插入的newNode是第一个节点
    if (this.length == 0) {
      this.head = newNode;
      this.tail = newNode;
    } else {
      // 原链表不为空
      // 情况2：position == 0
      if (position == 0) {
        this.head.prev = newNode;
        newNode.next = this.head;
        this.head = newNode;
      } else if (position == this.length) {
        // 情况3：position == this.length
        this.tail.next = newNode;
        newNode.prev = this.tail;
        this.tail = newNode;
      } else {
        // 情况4：0 < position < this.length
        let currentNode = this.head;
        let index = 0;

        while (index < position) {
          currentNode = currentNode.next;
          index++;
        }

        //修改pos位置前后节点变量的指向
        newNode.next = currentNode;
        newNode.prev = currentNode.pre;
        currentNode.prev.next = newNode;
        currentNode.prev = newNode;
      }
    }
    this.length++;
    return Status.OK;
  }
  indexOf(data) {
    let currentNode = this.head;
    let index = 0;
    while (currentNode) {
      if (currentNode.data === data) {
        return index;
      }
      currentNode = currentNode.next;
      index++;
    }
    return -1;
  }
  removeAt(position) {
    if (position < 0 || position >= this.length) return Status.ERROR;
    let currentNode = this.head;
    if (this.length === 1) {
      this.head = null;
      this.tail = null;
    } else {
      if (position === 0) {
        // 删除第一个节点
        this.head.next.prev = null;
        this.head = this.head.next;
      } else if (position === this.length - 1) {
        // 删除最后一个节点
        currentNode = this.tail;
        this.tail.prev.next = null;
        this.tail = this.tail.prev;
      } else {
        let index = 0;
        while (index < position) {
          currentNode = currentNode.next;
          index++;
        }
        currentNode.pre.next = currentNode.next;
        currentNode.next.pre = currentNode.pre;
      }
    }
    this.length--;
    return currentNode.data;
  }
  remove(data) {
    let index = this.indexOf(data);
    return this.removeAt(index);
  }
  getElem(position) {
    if (position < 0 || position >= this.length) {
      return Status.ERROR;
    }
    let currentNode = null;
    let index = 0;
    if (position < this.length / 2) {
      currentNode = this.head;
      while (index < position) {
        currentNode = currentNode.next;
        index++;
      }
    } else {
      currentNode = this.tail;
      index = this.length - 1;
      while (index > position) {
        currentNode = currentNode.prev;
        index--;
      }
    }
    return currentNode.data;
  }
}
let list = new doubleList();
list.insertElem(0, 'a');
list.insertElem(1, 'b');
list.remove('a');
let res = list.getElem(0);

console.log(list, res);
