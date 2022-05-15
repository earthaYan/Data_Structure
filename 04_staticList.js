const Status = {
  OK: 1,
  ERROR: 0,
  TRUE: 1,
  FALSE: 0,
};
class staticList {
  constructor(maxSize) {
    this.maxSize = maxSize;
    let space = new Array(maxSize);
    this.List = space;
  }
  initList() {
    const len = this.List.length;
    for (let i = 0; i < len - 1; i++) {
      this.List[i].data = i;
      this.List[i].cur = i + 1;
    }
    this.List[len - 1].cur = 0;
    return Status.OK;
  }
  malloc() {
    // 申请节点
    let i = this.List[0].cur;
    if (this.List[0].cur) {
      this.List[0].cur = this.List[i].cur;
    }
    return i;
  }
  insert(index, ele) {
    let maxLen = this.maxSize - 1;
    if (index < 1 || index > this.List.length + 1) {
      return Status.ERROR;
    }
    let numForFree = this.malloc();
    if (numForFree) {
      this.List[numForFree].data = ele;
      for (let i = 1; i < index - 1; i++) {
        maxLen = this.List[maxLen].cur;
      }
      this.List[numForFree].cur = this.List[maxLen].cur;
      this.List[maxLen].cur = numForFree;
      return Status.OK;
    }
    return Status.ERROR;
  }
  free(index) {
    this.List[index].cur = this.List[0].cur;
    this.List[0].cur = index;
  }
  del(index) {
    if (index < 1 || index > this.List.length) {
      return Status.ERROR;
    }
    let len = this.List.length - 1;
    for (let j = 1; j < index - 1; j++) {
      len = this.List[len].cur;
    }
    j = this.List[len].cur;
    this.List[len].cur = this.List[j].cur;
    this.free(j);
    return Status.OK;
  }
}
