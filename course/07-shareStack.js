Status = {
  OK: 1,
  ERROR: 0,
  TRUE: 1,
  FALSE: 0,
};
class shareStack {
  constructor(length) {
    this.stack = new Array(length);
    this.top1 = -1;
    this.top2 = this.stack.length;
    this.rightPositionId = 0;
  }
  sharePush(ele, stackId) {
    if (this.top2 - this.top1 === 1) {
      console.log('栈已满');
      return Status.ERROR;
    }
    if (stackId === 1) {
      this.top1++;
      this.stack[this.top1] = ele;
    } else if (stackId === 2) {
      this.top2--;
      this.stack[this.top2] = ele;
      this.rightPositionId = this.top2;
    }
    return Status.OK;
  }
  sharePop(stackId) {
    if (stackId === 1) {
      if (this.top1 === -1) {
        // 空栈
        return Status.ERROR;
      }
      const data = this.stack.splice(this.top1, 1);
      this.top1--;
      this.rightPositionId--;
      return data;
    } else if (stackId === 2) {
      if (this.top2 === this.stack.length) {
        return Status.ERROR;
      }
      const data = this.stack.splice(this.rightPositionId, 1);
      this.top2++;
      return data;
    }
  }
}

const newShare = new shareStack(5);
newShare.sharePush('a', 1);
newShare.sharePush('b', 2);
newShare.sharePush('c', 2);
newShare.sharePush('d', 2);
newShare.sharePush('e', 2);
newShare.sharePush('f', 1);
console.log(newShare);
newShare.sharePop(1);
console.log(newShare);
newShare.sharePop(2);
console.log(newShare);
