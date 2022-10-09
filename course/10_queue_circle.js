Status = {
  OK: 1,
  ERROR: 0,
  TRUE: 1,
  FALSE: 0,
};
class CircleQueue {
  constructor(maxSize) {
    this.queue = [];
    this.front = 0;
    this.rear = 0;
    this.maxSize = maxSize;
  }
  getLength() {
    // 返回队列的元素个数，即队列当前长度
    return (this.rear - this.front + this.maxSize) % this.maxSize;
  }
  EnQueue(ele) {
    // 若队列未满，则队尾插入元素
    if ((this.rear + 1) % this.maxSize === this.front) {
      return Status.ERROR;
    }
    this.queue[this.rear] = ele;
    // rear指针后移一位，若到最后则转到数组头部
    this.rear = (this.rear + 1) % this.maxSize;
    return Status.OK;
  }
  DeQueue() {
    if (this.front === this.rear) {
      return Status.ERROR;
    }
    const ele = this.queue(this.front);
    this.front = (this.front + 1) % this.maxSize;
    return Status.OK + ele;
  }
}
