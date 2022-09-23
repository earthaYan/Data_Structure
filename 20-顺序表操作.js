/**
 *
 * @param {Array} arr 待插入数组
 * @param {number} pos 指定位置
 * @param {number} x  待插入数字
 * @returns
 */
function insert(arr, pos, x) {
  if (pos === -1) {
    arr[arr.length] = x;
    return arr;
  }
  for (let j = arr.length - 1; j >= pos; j--) {
    arr[j + 1] = arr[j];
  }
  arr[pos] = x;
  return arr;
}

function remove(arr, pos) {
  if (pos === -1) {
    arr.length--;
    return arr;
  }
  for (let i = pos; i < arr.length - 1; i++) {
    arr[i] = arr[i + 1];
  }
  return arr;
}
function scanFromLeftToRight(arr, x) {
  for (let i = 0; i < arr.length; i++) {
    if (arr[i] === x) {
      return i;
    }
  }
  return -1;
}
function scanFromRightToLeft(arr, x) {
  for (let i = arr.length - 1; i > 0; i--) {
    if (arr[i] === x) {
      return i;
    }
  }
  return -1;
}
// 带表头监督元的顺序查找（从右向左）
function QSearch(arr, x) {
  let start = arr.length;
  let arrTemp = insert(arr, 0, x);
  while (arrTemp[start] !== x) {
    start--;
  }
  return start--; //返回真实位置
}
let arr1 = [3, 5, 10, 16, 7, 32, 83, 23];
QSearch(arr1, 3);
