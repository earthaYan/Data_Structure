// 目的：设计算法，将数组arr的每一个元素都循环的右移K位，0<k<n
/**
 * 自然语言描述算法：每遍将n个元素循环的右移一位，经过K遍操作而完成
 */
function moveRight1(arr, k) {
  for (let i = 0; i < k; i++) {
    const temp = arr[arr.length - 1];
    for (let j = arr.length - 2; j > 0; j--) {
      arr[0] = temp;
    }
  }
}
