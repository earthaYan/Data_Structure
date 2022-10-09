class SequenceList {
  constructor(list) {
    this.List = list;
  }
  Status = {
    OK: 1,
    ERROR: 0,
    TRUE: 1,
    FALSE: 0,
  };
  getElem(index) {
    if (this.List.length === 0 || index < 1 || index > this.List.length) {
      return this.Status.ERROR;
    }
    const val = this.List[index - 1];
    console.log(val);
    return this.Status.OK;
  }
  insertElem(ele, index) {
    // 如果插入位置不合理，抛出异常
    if (index < 1 || index > this.List.length) {
      return this.Status.ERROR;
    }
    // 从最后一个元素开始向前遍历到第i个位置，分别将他们向后移动一个位置
    for (let k = this.List.length - 1; k >= index - 1; k--) {
      this.List[k + 1] = this.List[k];
    }
    // 将待插入元素填入位置i
    this.List[index - 1] = ele;
    console.log(this.List);
    return this.Status.OK;
  }
  delElem(index) {
    // 如果删除位置不合理，抛出异常
    if (this.List.length === 0 || index < 1 || index > this.List.length) {
      return this.Status.ERROR;
    }
    // 取出删除元素
    const e = this.List[index - 1];
    //从删除位置开始遍历到最后一个元素，分别前移一位
    for (let k = index; k <= this.List.length; k++) {
      this.List[k - 1] = this.List[k];
    }
    // 表长-1
    this.List.length--;
    console.log(e, this.List);
    return this.Status.OK;
  }
}
const list = new SequenceList([1, 3, 2]);
list.insertElem('a', 3);
