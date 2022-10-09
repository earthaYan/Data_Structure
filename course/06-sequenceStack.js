// 栈的顺序存储
class SequenceStack {
  constructor() {
    this.length = 0;
    this.stack = [];
  }
  // 入栈push
  pushToStack(ele) {
    this.length++;
    this.stack[this.length - 1] = ele;
  }
  // 出栈pop
  popFromStack() {
    if (this.length === 0) {
      return 'error';
    }
    const index = this.length - 1;
    const data = this.stack[index];
    this.stack.length--;
    this.length--;
    return data;
  }
}
const newStack = new SequenceStack();
newStack.pushToStack('a');
newStack.pushToStack('b');
newStack.popFromStack();
console.log(newStack);
