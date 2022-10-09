class List {
  constructor(listA, listB) {
    this.listA = listA;
    this.listB = listB;
  }
  getLength(list) {
    return list.length;
  }
  findElem(list, val) {
    return list.findIndex((item) => item === val);
  }
  inertToTargetList(list, val) {
    let len = this.getLength(list);
    list[len] = val;
  }
  unionList() {
    let lenA = this.getLength(this.listA);
    // 遍历列表A
    for (let i = 0; i < lenA; i++) {
      // 获取到当前值
      const val = this.listA[i];
      // 如果当前值不存在于B中，则将该值插入到列表B中
      if (this.findElem(this.listB, val) === -1) {
        this.inertToTargetList(this.listB, val);
      }
    }
    return this.listB;
  }
}
const listA = [1, 2, 4];
const listB = [2, 3, 4, 5];
const obj = new List(listA, listB);
console.log(obj.unionList());
